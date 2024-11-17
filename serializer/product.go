package serializer

import "mall/model"

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	CategoryID   uint   `json:"category_id"`
	Title        string `json:"title"`
	Info         string `json:"info"`
	ImgPath      string `json:"img_path"`
	Price        string `json:"price"`
	DiscoutPrice string `json:"discount_price"`
	View         uint64 `json:"view"`
	CreatedAt    int64  `json:"created_at"`
	Num          int    `json:"num"`
	OnSale       bool   `json:"on_sale"`
	BossID       uint   `json:"boss_id"`
	BossName     string `json:"boss_name"`
	BossAvatar   string `json:"boss_avatar"`
}

// 序列化商品
func BuildProduct(item *model.Product) Product {
	return Product{
		ID:           item.ID,
		Name:         item.Name,
		CategoryID:   item.Category,
		Title:        item.Title,
		Info:         item.Info,
		ImgPath:      item.ImgPath,
		Price:        item.Price,
		DiscoutPrice: item.DiscountPrice,
		View:         item.View(),
		Num:          item.Num,
		OnSale:       item.OnSale,
		CreatedAt:    item.CreatedAt.Unix(),
		BossID:       item.BossId,
		BossName:     item.BossName,
		BossAvatar:   item.BossAvatar,
	}
}

func BuildProducts(items []*model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
