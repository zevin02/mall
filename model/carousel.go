package model

import "gorm.io/gorm"

// 商品列的轮廓图，里面包含的就是单凭的图片和产品id
type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductId uint `gorm:"not null"`
}
