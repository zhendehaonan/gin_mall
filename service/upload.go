package service

import (
	"gin_mall/conf"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

// 上传图片到本地
func UpdateAvatarToLocalStatic(file multipart.File, userId uint, username string) (filepath string, err error) {
	bId := strconv.Itoa(int(userId)) //id转为字符串类型   用作路径拼接
	bashPath := "." + conf.AvatarPAth + "user" + bId + "/"
	if !DirExistOrNot(bashPath) {
		CreateDir(bashPath)
	}
	avatarPath := bashPath + username + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bId + "/" + username + ".jpg", nil
}

// 上传图片到本地
func UploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filepath string, err error) {
	bId := strconv.Itoa(int(userId)) //id转为字符串类型   用作路径拼接
	bashPath := "." + conf.ProductPath + "boss" + bId + "/"
	if !DirExistOrNot(bashPath) {
		CreateDir(bashPath)
	}
	productPath := bashPath + productName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
}

// 判断该路径是否存在
func DirExistOrNot(filePath string) bool {
	stat, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}
