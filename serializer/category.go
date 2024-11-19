package serializer

import "mall/model"

// 这个就是存储图片,把这个进行返回给到用户
type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

// 序列化轮廓图
func BuildCategory(item model.Category) Category {
	return Category{
		ID:           item.ID,
		CreatedAt:    item.CreatedAt.Unix(),
		CategoryName: item.CategoryName,
	}
}

// 序列化轮廓图列表
func BuildCategorys(items []model.Category) (categorys []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categorys = append(categorys, category)
	}
	return categorys
}
