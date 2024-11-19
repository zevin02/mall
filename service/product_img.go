package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/serializer"
	"strconv"
)

type ListProductImg struct {
}

func (service *ListProductImg) List(ctx context.Context, pid string) serializer.Response {
	var products []*model.ProductImg //商品都是商家创建的

	product_id, _ := strconv.Atoi(pid)
	productImgDao := dao.NewProductImgDao(ctx)
	products, _ = productImgDao.ListProductImg(uint(product_id))
	return serializer.BuildListResponse(serializer.BUildProductImgs(products), uint(len(products)))
}
