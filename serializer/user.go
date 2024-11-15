package serializer

import (
	"mall/model"
	"mall/pkg/util"
)

// VO view object 前端查看的数据类型,相当于是给到前端的proto
// 每个字段都有一个标签，告诉go的json库，将结构体转化成json时使用的key名
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`
	Status   string `json:"status"`
	CreateAt int64  `json:"create_at"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

func BuildUser(user *model.User) *User {
	//把从数据库中的数据处理后给到前端
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,

		Status:   user.Status,
		CreateAt: user.CreatedAt.Unix(),
		Avatar:   user.Avatar,
		Email:    user.Email,
	}
}

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) *Money {
	util.Encrypt.SetKey(key)
	return &Money{
		UserId:    user.ID,
		UserName:  user.UserName,
		UserMoney: util.Encrypt.AesDecoding(user.Money), //对密码进行解密
	}
}
