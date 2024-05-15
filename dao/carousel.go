package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// 查找轮播图Carousel
func (carouselDao *CarouselDao) ListCarousel() (carousel []*model.Carousel, err error) {
	err = carouselDao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
