package conf

import (
	"gopkg.in/ini.v1"
	"mall/dao"
	"strings"
)

var (
	AppModel string
	HttpPort string

	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDB       string
	RedisAddr     string
	RedisPassword string
	RedisDbName   string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

// Init init config and db conncection
func Init() {
	//本地文件中读取环境变量
	file, err := ini.Load("./conf/config.ini") //加载配置文件
	if err != nil {
		panic(err)
	}
	//加载一下配置
	LoadServer(file)
	LoadMySql(file)
	LoadEmail(file)
	LoadRedis(file)
	LoadPhotoPath(file)

	//mysql的读写分离
	//mysql 读 主,构建用于读取数据的mysql连接的字符串
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

	//mysql 写 (主从复制)，构建用于写的myslq的字符串
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	//传入读写连接字符串以初始化数据库连接
	dao.Database(pathRead, pathWrite)

}

func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()

}

func LoadRedis(file *ini.File) {
	RedisDB = file.Section("redis").Key("RedisDB").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPassword = file.Section("redis").Key("RedisPassword").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()

}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()

}

func LoadMySql(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String() //获得service模块下的appmode的数据
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()

}

func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppMode").String() //获得service模块下的appmode的数据
	HttpPort = file.Section("service").Key("HttpPort").String()

}
