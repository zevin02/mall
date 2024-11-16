package service

import (
	"context"
	"gopkg.in/mail.v2"
	"mall/conf"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
	"mime/multipart"
	"strings"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //前端验证,相当于是支付密码，使用这个进行加解密
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	// 这个直接和notice中的id进行对应
	//1. 绑定邮箱
	//2. 解除绑定邮箱
	//3. 改密码
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

// 展示用户金额
func (service *ShowMoneyService) ShowMoney(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
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
		Data:   serializer.BuildMoney(user, service.Key),
	}

}

// 验证邮箱
func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS
	if token == "" {
		code = e.InvalidParams
	} else {
		//解析token
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else {
			//从token中解析出来email和处理操作
			userId = claims.UserID //提取当前是哪个用户
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}

	}
	if code != e.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//获取用户的信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if operationType == 1 {
		//绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		//解除绑定
		user.Email = ""
	} else if operationType == 3 {
		err = user.SetPassword(password)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
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

// 发送邮箱,这个id就是通过从登陆态的token中解析出来的
func (service *SendEmailService) Send(ctx context.Context, id uint) serializer.Response {
	code := e.SUCCESS
	var address string       //用于保存发送通知的邮箱地址
	var notice *model.Notice //绑定邮箱，修改密码，模板通知，用于邮件内容
	//生成用于邮箱验证或密码重置的令牌
	//根据当前的请求生成对应的token，之后通过检查当前的这个token来设置这个对应的操作
	token, err := util.GenerateEmailToken(id, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	noticeDao := dao.NewNoticeDao(ctx)
	//根据操作的类型，获取通知
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	address = conf.ValidEmail + token //发送方邮件地址和令牌

	mailStr := notice.Text //获取邮件模板文本
	//题换Email为实际的邮件地址
	mailTex := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail) //设置发送人的邮箱地址和
	m.SetHeader("To", service.Email)    //设置接收人的邮箱地址
	m.SetHeader("Subject", "FanOne")    //设置邮件的主题
	m.SetBody("text/html", mailTex)     //设置邮箱内容
	//创建smtp拨号其并配置
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	//尝试发送有哦间

	if err := d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

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

func (service *UserService) Post(ctx context.Context, id uint, file multipart.File, size int64) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(id) //根据主键id get user info
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//保存图片在本地
	path, err := UploadAvatarToLocalStatic(file, id, user.UserName)

	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path //把当前图片存储的路径存储起来就行
	err = userDao.UpdateUserById(id, user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}

}
