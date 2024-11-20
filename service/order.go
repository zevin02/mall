package service

import (
	"context"
	"fmt"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

// 一个订单中包含哪个产品，数量，地址，金额，商家，用户
type OrderService struct {
	ProductId uint `form:"product_id" json:"product_id"`
	Num       int  `form:"num" json:"num"`
	AddressId uint `form:"address_id" json:"address_id"`
	Money     int  `form:"money" json:"money"`
	BossId    uint `form:"boss_id" json:"boss_id"`
	UserId    uint `form:"user_id" json:"user_id"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, id uint) serializer.Response {
	var err error
	code := e.SUCCESS
	orderDao := dao.NewOrderDao(ctx)

	//order为空，说明没有收藏过
	order := &model.Orders{
		UserId: service.UserId,
		BOssId: service.BossId,
		Num:    service.Num,
		Type:   model.NOTPAY, //默认未支付当前订单
		Money:  float64(service.Money),
	}
	//检验地址是否存在

	number := fmt.Sprintf("09v", rand.New(rand.NewSource(time.Now().Unix())).Int31n(1000000))
	//更新当前的订单字符串
	number = number + strconv.Itoa(int(service.ProductId)) + strconv.Itoa(int(service.UserId))
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	err = orderDao.CreateOrder(order)
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
