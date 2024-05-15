package model

// 商品分类
import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string //商品类别
}
