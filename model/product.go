package model

import (
	"gorm.io/gorm"
	"mall/cache"
	"strconv"
)

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
	BossId        uint //当前商品是哪个商家创建的
	BossName      string
	BossAvatar    string
}

// 获取点击数
func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// 增加商品点击数
func (product *Product) IncrView() {
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID))) //将这个商品在每日商品排行榜中加一

}
