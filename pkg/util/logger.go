package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := setOutPutFile()

	if LogrusObj != nil {
		//如果已经实例化了，就不需要再实例化了
		LogrusObj.Out = src
		return
	}
	//实例化
	logger := logrus.New()
	logger.Out = src                   //设置输出
	logger.SetLevel(logrus.DebugLevel) //设置日志级别
	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
	}) //设置格式
	//加个hook形成elk体系，将日志写入到es中，方便后续对日志的查看
	//hook := model.EsHookLog()
	//logger.AddHook(hook)
	//后续的
	LogrusObj = logger
}

// 返回一个文件
func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	//获得当前的路径
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		//当前路径不存在，就创建
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log" //文件名
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if os.IsNotExist(err) {
		//当前路径不存在，就创建
		if err = os.MkdirAll(fileName, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend) //
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return src, nil

}
