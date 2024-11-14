package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 通过用户名判断用户是否存在
func (dao *UserDao) ExistOrNotByUserName(username string) (user *model.User, exist bool, err error) {
	//model指定要查看的模型，
	//find将查询的结果存储到user中
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", username).Find(&user).Count(&count).Error
	//如果没有找到用户或者用户为空
	if count == 0 {
		//count=0说明没有找到
		return nil, false, nil
	}
	//find it,把查询到的数据进行返回
	return user, true, err

}

// CreateUser 创建一个用户,传入指针就不用发生一个拷贝
func (dao *UserDao) CreateUser(user *model.User) error {
	//直接这样创建即可
	return dao.DB.Model(&model.User{}).Create(&user).Error

}

func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	//First 用于检索满足条件的第一条记录
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return user, err
}

func (dao *UserDao) UpdateUserById(id uint, user *model.User) error {
	err := dao.Model(&model.User{}).Where("id=?", id).Updates(user).Error
	return err
}
