package api

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	addressService := service.AddressService{}
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.Create(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user CreateAddress api err:", err)
	}
}

func DeleteAddress(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	addressService := service.AddressService{}
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.Delete(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user DeleteAddress api err:", err)
	}
}

func UpdateAddress(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	addressService := service.AddressService{}
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.Update(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user UpdateAddress api err:", err)
	}
}

func ShowAddress(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	addressService := service.AddressService{}
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.Show(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user ShowAddress api err:", err)
	}
}
