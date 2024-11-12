package model

import "gorm.io/gorm"

// 收藏夹
type Favorite struct {
	gorm.Model
	User   User `gorm:"ForeignKey:UserId"` //设置外键，通过userid来进行关联
	UserId uint `gorm:"not null"`

	Product   Product `gorm:"ForeignKey:ProductId"` //设置外键，通过userid来进行关联
	ProductId uint    `gorm:"not null"`

	Boss   User `gorm:"ForeignKey:BossId"` //商家,通过bossid来连接
	BossId uint `gorm:"not null"`
}
