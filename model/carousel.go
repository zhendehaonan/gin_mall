package model

// 轮播图
import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string //图片路径
	ProductId uint   `gorm:"not null"` //图片id
}
