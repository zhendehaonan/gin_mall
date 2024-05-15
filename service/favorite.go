package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
	"strconv"
)

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (service *FavoriteService) List(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	favorite, err := favoriteDao.ListFavorite(uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "获取失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorite), uint(len(favorite)))
}

func (service *FavoriteService) Create(ctx context.Context, uid uint) serializer.Response {
	var err error
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	//判断商品是否被收藏
	exist, err := favoriteDao.IsOrNotExist(service.ProductId, uid)
	//已存在
	if exist {
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg:    "商品已被收藏",
			Error:  err.Error(),
		}
	}
	//不存在
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	favorite := &model.Favorite{
		User:      *user,
		UserId:    uid,
		Product:   product,
		ProductId: service.ProductId,
		Boss:      *boss,
		BossId:    service.BossId,
	}
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "收藏失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "收藏成功",
	}
}

func (service *FavoriteService) Delete(ctx context.Context, uid uint, fId string) serializer.Response {
	fid, _ := strconv.ParseUint(fId, 10, 64)
	var err error
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	err = favoriteDao.DeleteFavorite(uid, uint(fid))
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
