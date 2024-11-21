package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type SkillGoodsDao struct {
	*gorm.DB
}

func NewSkillGoodsDao(ctx context.Context) *SkillGoodsDao {
	return &SkillGoodsDao{NewDBClient(ctx)}
}

func (dao *SkillGoodsDao) Create(in *model.SkillGoods) error {
	return dao.Model(&model.SkillGoods{}).Create(&in).Error
}

// 一次性插入多条记录
func (dao *SkillGoodsDao) CreateByList(in []*model.SkillGoods) error {
	return dao.Model(&model.SkillGoods{}).Create(&in).Error
}

// 查询当前秒杀活动中还有剩余数量的 skill good
func (dao *SkillGoodsDao) ListSkillGoods() (resp []*model.SkillGoods, err error) {
	err = dao.Model(&model.SkillGoods{}).Where("num > 0").Find(&resp).Error
	return
}
