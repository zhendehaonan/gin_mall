package model

// 购物车
import "gorm.io/gorm"

// 购物车
type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"` //用户id
	ProductId uint `gorm:"not null"` //商品id
	BossId    uint `gorm:"not null"` //商家id
	Num       uint `gorm:"noy null"` //数量
	MaxNum    uint `gorm:"noy null"` //最大购买数量
	Check     bool //是否支付
}
