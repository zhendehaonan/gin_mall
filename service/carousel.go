package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
)

type CarouselService struct {
}

func (carouselService *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.Success
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "获取轮播图失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
