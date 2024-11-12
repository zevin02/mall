package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"` //用户名 设置为唯一键
	Email          string
	PasswordDigest string
	NickName       string
	Status         string //用户状态，是否被封禁
	Avatar         string
	Money          string //钱是密文，所以用string
}
