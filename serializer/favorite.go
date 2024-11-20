package serializer

import (
	"context"
	"mall/dao"
	"mall/model"
)

// 这个就是存储图片,把这个进行返回给到用户
type Favorite struct {
	userId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	Category      uint   `json:"category"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

// 序列化轮廓图
func BuildFavorite(item *model.Favorite, product *model.Product) Favorite {
	return Favorite{
		userId:        item.UserId,
		ProductId:     item.ProductId,
		CreatedAt:     item.CreatedAt.Unix(),
		Name:          product.Name,
		Category:      product.Category,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		BossId:        item.BossId,
		Num:           product.Num,
		OnSale:        product.OnSale,
	}
}

// 序列化轮廓图列表
func BuildFavorites(ctx context.Context, items []*model.Favorite) (favorites []Favorite) {
	productDao := dao.NewProductDao(ctx)

	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}

		favorite := BuildFavorite(item, product)
		favorites = append(favorites, favorite)
	}
	return favorites
}
