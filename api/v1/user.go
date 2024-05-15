package api

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册
// 这里获取参数，可以理解为前端发来的参数都在ctx中
func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService
	if err := ctx.ShouldBind(&userRegister); err == nil { //ShouldBond绑定数据到userRegister中
		res := userRegister.Register(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user register api err:", err)
	}
}

// 用户登录
// 这里获取参数，可以理解为前端发来的参数都在ctx中
func UserLogin(ctx *gin.Context) {
	var userLogin service.UserService
	if err := ctx.ShouldBind(&userLogin); err == nil { //ShouldBond绑定数据到userLogin中
		res := userLogin.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user login api err:", err)
	}
}

// 用户修改信息(修改用户名username)
// 这里获取参数，可以理解为前端发来的参数都在ctx中
func UserUpdate(ctx *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization")) //解析token
	if err := ctx.ShouldBind(&userUpdate); err == nil {          //ShouldBond绑定数据到userUpdater中
		res := userUpdate.Update(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user update api err:", err)
	}
}

// 修改头像
// 这里获取参数，可以理解为前端发来的参数都在ctx中
func UpdateAvatar(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")
	fileSize := fileHeader.Size
	var UpdateAvatar service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization")) //解析token
	if err := ctx.ShouldBind(&UpdateAvatar); err == nil {        //ShouldBond绑定数据到UpdateAvatar中
		res := UpdateAvatar.Post(ctx.Request.Context(), claims.ID, file, fileSize)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user updateAvatar api err:", err)
	}
}

// 发送邮箱
func SendEmail(ctx *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization")) //解析token
	if err := ctx.ShouldBind(&sendEmail); err == nil {           //ShouldBond绑定数据到sendEmail中
		res := sendEmail.Send(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user sendEmail api err:", err)
	}
}

// 验证邮箱
func ValidEmail(ctx *gin.Context) {
	var validEmail service.ValidEmailService
	if err := ctx.ShouldBind(&validEmail); err == nil { //ShouldBond绑定数据到validEmail中
		res := validEmail.Valid(ctx.Request.Context(), ctx.GetHeader("Authorization"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user validEmail api err:", err)
	}
}

// ShowMoney 显示金额
func ShowMoney(ctx *gin.Context) {
	var showMoney service.ShowMoneyService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization")) //解析token
	if err := ctx.ShouldBind(&showMoney); err == nil {           //ShouldBond绑定数据到showMoney中
		res := showMoney.Show(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user ShowMoney api err:", err)
	}
}
