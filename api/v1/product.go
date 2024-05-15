package api

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建商品
func CreateProduct(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := ctx.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(ctx.Request.Context(), claim.ID, files)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user CreateProduct api err:", err)
	}
}

// 分页查询商品
func ListProduct(ctx *gin.Context) {
	listProductService := service.ProductService{}
	if err := ctx.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user listProduct api err:", err)
	}
}

// 搜索商品
func SearchProduct(ctx *gin.Context) {
	listProductService := service.ProductService{}
	if err := ctx.ShouldBind(&listProductService); err == nil {
		res := listProductService.Search(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user searchProduct api err:", err)
	}
}

// 商品信息展示
func ShowProduct(ctx *gin.Context) {
	showProductService := service.ProductService{}
	if err := ctx.ShouldBind(&showProductService); err == nil {
		res := showProductService.Show(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user showProduct api err:", err)
	}
}
