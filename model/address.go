package model

import "gorm.io/gorm"

// Address 通常用于与数据库交互，特别是使用GORM这个ORM库
// 用户的地址
type Address struct {
	//可以直接使用address访问model中的字段
	gorm.Model //嵌入时的结构体，包含了一些常用的字段，ID，CREATEAT，UPDATEAT，deleteat，简化了对模型的定义
	//结构体标签，为了结构体字段提供额外信息的字符串，和其他框架进行交互
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Address string `gorm:"type:varchar(50) not null"`
}
