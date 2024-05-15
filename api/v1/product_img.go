package api

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProductImg(ctx *gin.Context) {
	listProductImgService := service.ProductImgService{}
	if err := ctx.ShouldBind(&listProductImgService); err == nil {
		res := listProductImgService.List(ctx.Request.Context(), ctx.Param("product_id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user listProductImg api err:", err)
	}
}
