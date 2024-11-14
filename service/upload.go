package service

import (
	"io/ioutil"
	"mall/conf"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, id uint, userName string) (string, error) {
	bId := strconv.Itoa(int(id)) //路径拼接
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := ioutil.ReadAll(file) //read file from http to content
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666) //write content to file

	if err != nil {
		return "", err
	}

	return "user" + bId + "/" + userName + ".jpg", nil

}

//判断文件夹路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()

}

//创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}
