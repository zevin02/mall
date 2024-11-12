package model

import "gorm.io/gorm"

// 商品模型
type Product struct {
	gorm.Model
	Name          string
	Category      uint //商品类型
	Title         string
	Info          string
	Price         string
	ImgPath       string
	DiscountPrice string
	OnSale        bool `gorm:"default false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}
