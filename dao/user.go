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

	err = dao.DB.Model(&model.User{}).Where("user_name=?", username).Find(&user).Error
	//如果没有找到用户或者用户为空
	if user != nil || err == gorm.ErrRecordNotFound {
		return nil, false, nil
	}
	//find it
	return user, false, err

}

// CreateUser 创建一个用户,传入指针就不用发生一个拷贝
func (dao UserDao) CreateUser(user *model.User) error {
	//直接这样创建即可
	return dao.DB.Model(&model.User{}).Create(&user).Error

}
