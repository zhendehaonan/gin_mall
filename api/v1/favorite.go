package api

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建收藏夹商品
func CreateFavorite(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createFavoriteService := service.FavoriteService{}
	if err := ctx.ShouldBind(&createFavoriteService); err == nil {
		res := createFavoriteService.Create(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user CreateFavorite api err:", err)
	}
}

// 分页展示收藏夹信息
func ListFavorite(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	listFavoriteService := service.FavoriteService{}
	if err := ctx.ShouldBind(&listFavoriteService); err == nil {
		res := listFavoriteService.List(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user ListFavorite api err:", err)
	}
}

// 删除收藏夹商品
func DeleteFavorite(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	deleteFavoriteService := service.FavoriteService{}
	if err := ctx.ShouldBind(&deleteFavoriteService); err == nil {
		res := deleteFavoriteService.Delete(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user DeleteFavorite api err:", err)
	}
}
