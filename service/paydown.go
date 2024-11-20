package service

import (
	"context"
	"fmt"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
	"strconv"
)

type OrderPay struct {
	OrderId uint   `json:"order_id" form:"order_id"`
	Money   int    `json:"money" form:"money"`
	OrderNo string `json:"order_no" form:"order_no"`
	PayTime string `json:"pay_time" form:"pay_time"`
	Num     int    `json:"num" form:"num"`
	Key     string `json:"key" form:"key"`   //支付密码
	Sign    string `json:"sign" form:"sign"` //签名

}

func (service *OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.SUCCESS
	orderdao := dao.NewOrderDao(ctx)
	tx := orderdao.Begin()                               //开启一个事务
	order, err := orderdao.GetOrderById(service.OrderId) //获取订单
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num) //计算总金额
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId) //获得到当前的用户信息

	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//对钱进行解密，减去订单，再加密保存
	moneyStr := util.Encrypt.AesDecoding(user.Money) //把当前用户的金额取出来
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	if moneyFloat-money < 0.0 {
		//用户总金额不足，就需要进行事务的回滚
		tx.Rollback()
		util.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	finMoney := fmt.Sprintf("%f", moneyFloat-money) //用户当前剩余的金额
	user.Money = util.Encrypt.AesEncoding(finMoney)
	userDao = dao.NewUserDaoByDB(userDao.DB)
	err = userDao.UpdateUserById(uId, user) //更新当前的用户信息，金额
	if err != nil {
		tx.Rollback()

		util.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	var boss *model.User
	boss, err = userDao.GetUserById(order.BOssId) //获取到商家的信息
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//商家的钱增加
	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)
	err = userDao.UpdateUserById(order.BOssId, boss)
	if err != nil {
		//失败就要进行事务的回滚
		tx.Rollback()

		util.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//商品数目减少
	var product *model.Product
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(order.ProductId)
	product.Num -= num
	err = productDao.UpdateProductById(order.ProductId, product)
	//更新订单状态
	order.Type = model.PAYED
	err = orderdao.UpdateOrderById(order.ID, order)

	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
