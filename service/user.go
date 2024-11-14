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
func (service *UserService) Register(ctx context.Context) serializer.Response {
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

func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	//判断用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist || err != nil {
		code = e.ErrorExistUserName
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册",
		}
	}
	//验证密码
	if user.CheckPassWord(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登陆",
		}
	}
	//每次登陆都要把token(携带一个认证的东西)分发给到用户，http是一个无状态的，一次请求，不会保存后续的状态
	//toekn的签发
	token, err := util.GenerateToken(user.ID, user.UserName, 0) //普通人设置权限为0
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登陆",
		}
	}
	//shioken
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}

}

// 修改用户信息
func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User

	code := e.SUCCESS
	//找到这个用户
	userDao := dao.NewUserDao(ctx)
	var err error
	user, err = userDao.GetUserById(uId)
	//uodate nickname
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}

}
