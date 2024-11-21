package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"mall/cache"
	"mall/dao"
	"mall/model"
	"mall/mq"
	"mall/pkg/util"
	"mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

type SkillGoodsService struct {
	SkillGoodsId uint   `json:"skill_goods_id" form:"skill_goods_id"` //秒杀id
	ProductId    uint   `json:"product_id" form:"product_id"`         //商品id
	BossId       uint   `json:"boss_id" form:"boss_id"`
	AddressId    uint   `json:"address_id" form:"address_id"`
	Key          string `json:"key" form:"key"`
}

// 将mysql中的秒杀商品导入到redis中
func (service *SkillGoodsService) InitSkillGoods(ctx context.Context) interface{} {
	skillGoodsDao := dao.NewSkillGoodsDao(ctx)
	skillGoods, _ := skillGoodsDao.ListSkillGoods() //查询所有可用的秒杀商品
	r := cache.RedisClient
	//加载到redis中
	for i := range skillGoods {
		util.LogrusObj.Info(*skillGoods[i])
		//给SK+ id这个key创建一个hash，两个field分别是num和money,将mysql中的数据导入到redis中
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "num", skillGoods[i].Num)
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "money", skillGoods[i].Money)

	}
	return nil
}

// 完成对某个秒杀单品的秒杀操作
func (service *SkillGoodsService) SkillGoods(ctx context.Context, id uint) serializer.Response {
	//从redis中获取当前秒杀商品的金额
	money, _ := cache.RedisClient.HGet("SK"+strconv.Itoa(int(service.SkillGoodsId)), "money").Float64() //获取当前需要秒杀的秒杀金额
	sk := &model.SkillGood2MQ{
		SkillGoodId: service.SkillGoodsId,
		ProductId:   service.ProductId,
		BossId:      service.BossId,
		UserId:      id,
		Money:       money,
		AddressId:   service.AddressId,
		Key:         service.Key,
	}
	//因为当前是高并发操作，所以需要保证分布式线程安全
	err := RedissionSecSkillGoods(sk)
	if err != nil {
		return serializer.Response{}

	}
	return serializer.Response{}

}

// 加锁
func RedissionSecSkillGoods(sk *model.SkillGood2MQ) error {
	p := strconv.Itoa(int(sk.ProductId)) //根据商品id来提取
	uuid := getUuid(p)
	_, err := cache.RedisClient.Del(p).Result() //先提前将这个uuid correspongding key delete
	//key是这把锁对应的名字，value correspond to this lock's client

	lockSucess, err := cache.RedisClient.SetNX(p, uuid, time.Second*3).Result()
	if err != nil || !lockSucess {
		//means get lock fail
		util.LogrusObj.Infoln("get lock fail", err)
		return errors.New("get lock fail")
	} else {
		util.LogrusObj.Infoln("get lock success")
	}
	//get lock success
	_ = SendSecSkillSToMQ(sk) //将数据发布到mq中进行消费
	//处理完之后，重新获得锁
	value, _ := cache.RedisClient.Get(p).Result()
	if value == uuid {
		//对比看当前锁是否被更换，相等就可以删除这把锁了
		_, err := cache.RedisClient.Del(p).Result()
		if err != nil {
			util.LogrusObj.Info("unlock fail")
			return nil
		} else {
			util.LogrusObj.Info("unlock success")
		}
	}
	return nil

}

func SendSecSkillSToMQ(sk *model.SkillGood2MQ) error {
	ch, err := mq.RabbitMQ.Channel()
	if err != nil {
		util.LogrusObj.Error(err)
		return err
	}
	//声明一个rabbitmq中的一个队列
	//mq重启后仍然存在，队列智能被当前连接使用，具有排他性能
	//队列是否不适用的时候自动删除，队列中创建一个非持久化的消息
	//声明一个名为 skill_goods 的持久化队列，如果队列已经存在，则不会重新创建。
	q, err := ch.QueueDeclare("skill_goods", true, false, false, false, nil)
	if err != nil {
		err = errors.New("rabbitMQ err:" + err.Error())
		return err
	}
	body, _ := json.Marshal(sk) //将当前的sk进行json化
	//将持久化的JSON格式的消息发布到skill_goods队列中

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbitMQ err:" + err.Error())
		return err
	}
	util.LogrusObj.Info(body)
	return nil
}

// 根据gid获取一个唯一的uid
func getUuid(gid string) string {
	codeLen := 8
	// 1. 定义原始字符串
	rawStr := "jkwangagDGFHGSERKILMJHSNOPQR546413890_"
	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	// 随机从中获取
	rand.Seed(time.Now().UnixNano())
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		b.WriteByte(rawStr[randNum])
	}
	return b.String() + gid
}

// mq的消费者一端
func MQ2MySQL() error {
	//redis
	r := cache.RedisClient
	ch, err := mq.RabbitMQ.Channel() //打开当前的mq的channel
	if err != nil {
		panic(err)
	}
	//确保skill_goods队列的存在
	q, err := ch.QueueDeclare("skill_goods", true, false, false, false, nil)
	_ = ch.Qos(1, 0, false) //消费者设置了通道的服务质量，确保消费者一次只处理一条消息
	msgs, _ := ch.Consume(q.Name, "", false, false, false, false, nil)
	//消费者不断等待数据到来
	for d := range msgs {
		//这个d就是消费测获得的消息
		//消费者开始消费
		var p model.SkillGood2MQ //这个是消费测获得的数据
		if err := json.Unmarshal(d.Body, &p); err != nil {
			//进行解码，转化成对应的结构化的数据
			d.Nack(false, false) //处理失败，拒绝并不重新入队列
			continue
		}
		//创建订单
		//订单扣除
		//redis扣除库存
		r.HIncrBy(strconv.Itoa(int(p.SkillGoodId)), "num", -1) //数量-1

		util.LogrusObj.Info("Done")
		d.Ack(false) //消息处理完之后，对当前消息确认ack
	}

	return nil
}
