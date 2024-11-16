package serializer

import "mall/model"

// 这个就是存储图片,把这个进行返回给到用户
type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
	CreatedAt int64  `json:"created_at"`
}

// 序列化轮廓图
func BuildCarousel(item model.Carousel) Carousel {
	return Carousel{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		ProductId: item.ProductId,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// 序列化轮廓图列表
func BuildCarousels(items []model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
