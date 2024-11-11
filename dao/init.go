package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	var db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead, //数据源名称（Data Source Name）用于指定连接数据库的详细数据
		DefaultStringSize:         256,      // string类型字段的默认长度
		DisableDatetimePrecision:  true,     // 精致datetime精度
		DontSupportRenameIndex:    true,     // 重命名索引
		DontSupportRenameColumn:   true,     // 用change重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger, //
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //命名策略不要加s
		},
	})

}
