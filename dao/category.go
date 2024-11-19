package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

// 商品列的轮廓图
type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

func (dao *CategoryDao) ListCategory() (category []model.Category, err error) {
	//找到全部符合条件的元素
	err = dao.DB.Model(&model.Category{}).Find(&category).Error
	return category, err
}
