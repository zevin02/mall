package model

import "gorm.io/gorm"

const (
	NOTPAY = 1
	PAYED  = 2
)

// 订单模型
type Orders struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	BOssId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	Num       int
	OrderNum  uint64  //订单数量
	Type      uint    //1 未支付，2已支付
	Money     float64 //当前订单多少钱，明文

}
