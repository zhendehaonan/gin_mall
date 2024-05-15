package dao

//数据模型迁移    也就是将数据模型迁移为数据库中的表
import (
	"fmt"
	"gin_mall/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Address{},
		&model.Admin{},
		&model.Carousel{},
		&model.Cart{},
		&model.Category{},
		&model.Favorite{},
		&model.Notice{},
		&model.Order{},
		&model.Product{},
		&model.ProductImg{},
		&model.User{},
	)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		return
	}
}
