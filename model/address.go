package model

//地址
import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId  uint   `gorm:"not null"`                  //用户id
	Name    string `gorm:"type:varchar(20) not null"` //用户名
	Phone   string `gorm:"type:varchar(11) not null"` //用户手机号
	Address string `gorm:"type:varchar(50) not null"` //用户地址
}
