package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) GetOrderById(id uint) (order *model.Orders, err error) {
	//First 用于检索满足条件的第一条记录
	err = dao.DB.Model(&model.Orders{}).Where("id=?", id).First(&order).Error
	return order, err
}

// 创建商品
func (dao *OrderDao) CreateOrder(order *model.Orders) error {
	//直接这样创建即可
	return dao.DB.Model(&model.Orders{}).Create(&order).Error

}

func (dao *OrderDao) UpdateOrderById(id uint, order *model.Orders) error {
	err := dao.Model(&model.Orders{}).Where("id=?", id).Updates(order).Error
	return err
}
