package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	//直接这样创建即可
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error

}
