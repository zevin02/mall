package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func (dao *FavoriteDao) GetFavoriteByUserIdProductId(userId, productId uint) (favorite model.Favorite, err error) {
	err = dao.DB.Where("user_id = ? AND product_id = ?", userId, productId).Find(&favorite).Error
	return favorite, err
}

// 创建商品
func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) error {
	//直接这样创建即可
	return dao.DB.Model(&model.Favorite{}).Create(&favorite).Error

}
