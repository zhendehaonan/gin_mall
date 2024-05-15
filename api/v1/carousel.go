package api

import (
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 轮播图
func ListCarousel(ctx *gin.Context) {
	var listCarousel service.CarouselService
	if err := ctx.ShouldBind(&listCarousel); err == nil { //ShouldBond绑定数据
		res := listCarousel.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
