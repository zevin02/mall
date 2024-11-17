package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

func (dao *ProductDao) GetNoticeById(id uint) (notice *model.Product, err error) {
	//First 用于检索满足条件的第一条记录
	err = dao.DB.Model(&model.Product{}).Where("id=?", id).First(&notice).Error
	return notice, err
}

// 创建商品
func (dao *ProductDao) CreateProduct(product *model.Product) error {
	//直接这样创建即可
	return dao.DB.Model(&model.Product{}).Create(&product).Error

}
