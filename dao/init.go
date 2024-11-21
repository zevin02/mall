package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var (
	_db *gorm.DB //gorm的全局变量
)

// 初始化mysql
func InitMySQL(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	//通过gorm简化对数据库的使用
	var db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead, //数据源名称（Data Source Name）用于指定连接数据库的详细数据
		DefaultStringSize:         256,      // string类型字段的默认长度
		DisableDatetimePrecision:  true,     // 精致datetime精度
		DontSupportRenameIndex:    true,     // 重命名索引
		DontSupportRenameColumn:   true,     // 用change重命名列
		SkipInitializeWithVersion: false,    //不跳过版本初始化
	}), &gorm.Config{
		//对gorm进行配置
		Logger: ormLogger, //设置gorm的日志记录其
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //命名策略不要加s
		},
	})
	if err != nil {
		return
	}
	//获取原始数据库对象以进行进一步的配置
	sqlDB, _ := db.DB()                        //这个就是对mysql的操作
	sqlDB.SetMaxIdleConns(20)                  //mysql的连接池,空闲
	sqlDB.SetMaxOpenConns(100)                 //打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) //一个连接的最大生命周期为 30 秒，这样可以确保连接在使用一段时间后被关闭，避免长时间占用。

	//主从配置
	_db = db //将db复制给全局变量_db
	//注册主从数据库配置，配置读写分离
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},                        //写操作连接，指向主库
		Replicas: []gorm.Dialector{mysql.Open(connWrite), mysql.Open(connWrite)}, //读操作，指向从库，这里使用两个从库来进行负载均衡
		Policy:   dbresolver.RandomPolicy{},                                      //随机策略选择读库
	}))

	migration() //进行数据库的迁移

}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
