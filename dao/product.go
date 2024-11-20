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

func (dao *ProductDao) CountProductByCond(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return total, err
}

func (dao *ProductDao) ListProductByCond(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Limit(page.PageSize).Offset(page.PageSize * (page.PageNum - 1)).Find(&products).Error
	return

}

func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, count int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").Count(&count).Error
	if err != nil {
		return
	}

	err = dao.DB.Model(&model.Product{}).Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Limit(page.PageSize).
		Offset(page.PageSize * (page.PageNum - 1)).
		Find(&products).Error
	return
}

func (dao *ProductDao) GetProductById(id uint) (user *model.Product, err error) {
	//First 用于检索满足条件的第一条记录
	err = dao.DB.Model(&model.Product{}).Where("id=?", id).First(&user).Error
	return user, err
}

func (dao *ProductDao) UpdateProductById(id uint, product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Where("id=?", id).Updates(product).Error
}
