package model

import "gorm.io/gorm"

//购物车模型
type Cart struct {
	gorm.Model //嵌入时的结构体，包含了一些常用的字段，ID，CREATEAT，UPDATEAT，deleteat，简化了对模型的定义
	//结构体标签，为了结构体字段提供额外信息的字符串，和其他框架进行交互
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	BossId    uint `gorm:"not null"` //商家id
	Num       uint `gorm:"not null"` //商品数量
	MaxNum    uint `gorm:"not null"` //商品限额
	Check     bool //是否支付

}