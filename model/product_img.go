package model

//商品图片
import "github.com/jinzhu/gorm"

type ProductImg struct {
	gorm.Model
	ProductID uint `gorm:"not null"`
	ImgPath   string
}
