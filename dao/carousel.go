package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

// 商品列的轮廓图
type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	//找到全部符合条件的元素
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return carousel, err
}
