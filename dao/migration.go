package dao

import (
	"fmt"
	"mall/model"
)

// Migration 自动迁移
func migration() {
	//设置表选项，包括字符集以支持更光放的字符
	//根据定义的模型结构体创建数据库表结构
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Admin{},
		&model.Category{},
		&model.Carousel{},
		&model.Cart{},
		&model.Notice{},
		&model.Product{},
		&model.ProductImg{},
		&model.Order{},
		&model.Favorite{},
	)

	if err != nil {
		fmt.Println("err", err)
	}
	return
}
