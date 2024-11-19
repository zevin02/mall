package serializer

import "mall/model"

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BUildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductId,
		ImgPath:   item.ImgPath,
	}
}

func BUildProductImgs(items []*model.ProductImg) (productimgs []ProductImg) {
	for _, item := range items {
		product := BUildProductImg(item)
		productimgs = append(productimgs, product)
	}
	return productimgs
}
