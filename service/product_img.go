package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"strconv"
)

type ProductImgService struct{}

func (s ProductImgService) List(ctx context.Context, productId string) serializer.Response {
	uid, _ := strconv.Atoi(productId)
	imgDao := dao.NewProductImgDao(ctx)
	imgs, err := imgDao.ListProductImg(uint(uid))
	if err != nil {
		return serializer.Response{
			Status: e.Error,
			Msg:    "获取商品图片失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildProductImgRes(imgs), uint(len(imgs))),
	}
}
