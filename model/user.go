package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 用户
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string
}

const (
	PasswordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

// 密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 校验密码
func (user User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	fmt.Println(err, "77777777777777777")
	return err == nil
}
