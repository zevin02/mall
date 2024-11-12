package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model     //嵌入时的结构体，包含了一些常用的字段，ID，CREATEAT，UPDATEAT，deleteat，简化了对模型的定义
	UserName       string
	PasswordDigest string
	Avatar         string
}
