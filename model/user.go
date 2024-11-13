package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// mysql中的user表
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"` //用户名 设置为唯一键
	Email          string
	PasswordDigest string //
	NickName       string
	Status         string //用户状态，是否被封禁
	Avatar         string
	Money          string //钱是密文，所以用string
}

const (
	PasswordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户

)

// SetPassword 传入的是原始的密码，表中存储的是用户的加密后的密码
func (user User) SetPassword(password string) error {
	//生成密码的摘要
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}
