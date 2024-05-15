package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
)

type CategoryService struct{}

func (service *CategoryService) List(ctx context.Context) interface{} {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	category, err := categoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "获取失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategories(category), uint(len(category)))
}
