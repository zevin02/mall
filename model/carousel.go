package model

import "gorm.io/gorm"

// 商品模型
type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductId uint `gorm:"not null"`
}
