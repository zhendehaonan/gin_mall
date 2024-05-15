package model

//管理员
import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName string //管理员名
	Password string //密码
	Avatar   string //头像
}
