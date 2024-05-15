package api

import (
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategories(ctx *gin.Context) {
	var listCategory service.CategoryService
	if err := ctx.ShouldBind(&listCategory); err == nil { //ShouldBond绑定数据
		res := listCategory.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
