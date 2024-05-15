package routes

import (
	"gin_mall/api/v1"
	"gin_mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})
		//用户注册
		v1.POST("/user/register", api.UserRegister)
		//用户登录
		v1.POST("/user/login", api.UserLogin)
		//轮播图
		v1.GET("/carousels", api.ListCarousel)
		//商品分页
		v1.GET("/products", api.ListProduct)
		//商品信息展示(某一个商品)
		v1.GET("/product/:id", api.ShowProduct)
		//商品图片信息展示(某一个商品)
		v1.GET("/imgs/:product_id", api.ListProductImg)
		//商品分类信息(所有商品)
		v1.GET("/categories", api.ListCategories)
		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			//用户修改信息
			//修改用户名
			authed.PUT("user", api.UserUpdate)
			//修改头像avatar
			authed.POST("avatar", api.UpdateAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)
			//显示金额
			authed.POST("money", api.ShowMoney)
			//创建商品
			authed.POST("product", api.CreateProduct)
			//搜索商品
			authed.POST("product/search", api.SearchProduct)

			//展示收藏夹
			authed.GET("favorites", api.ListFavorite)
			//添加商品到收藏夹
			authed.POST("favorites", api.CreateFavorite)
			//删除收藏夹中的某一个商品
			authed.DELETE("favorites/:id", api.DeleteFavorite)
			//创建地址
			authed.POST("address", api.CreateAddress)
			//展示某一个地址
			authed.GET("address/:id", api.ShowAddress)
			//展示所有地址
			//	authed.GET("address", api.ListAddress)
			//更新地址
			authed.PUT("address/:id", api.UpdateAddress)
			//删除地址
			authed.DELETE("address/:id", api.DeleteAddress)
		}
	}
	return r
}
