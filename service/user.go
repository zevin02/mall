package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //前端验证
}

// Register 用户注册函数,里面包含了从gin里面提取的请求参数
func (service UserService) Register(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	//密钥
	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}

	}
	//100000  ---->密文存储 ,对称加密操作
	util.Encrypt.SetKey(service.Key) //设置密钥

	userDao := dao.NewUserDao(ctx) //创建一个用户的dao

	_, exist, err := userDao.ExistOrNotByUserName(service.UserName) //传入请求的参数username
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		//用户已经exist
		code = e.ErrorExistName
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//用户不存在，可以注册
	user := model.User{
		UserName: service.UserName,
		Email:    "",
		NickName: service.NickName,
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    util.Encrypt.AesEncoding("100000"), //给每个用户初始100000元,加密,使用传入的密钥进行加密

	}
	//user表中设置密码,密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncoding
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.ERROR

	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
