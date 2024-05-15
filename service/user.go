package service

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/dto"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

// 前端传过来的数据格式（json）
type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //密钥
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"` //1:绑定邮箱 2：解绑邮箱 3：修改密码
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

// 用户注册
func (userService *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if userService.Key == "" || len(userService.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥不满足要求",
		}
	}
	//用户起始金额为10000，通过密钥加密后存储到数据库中
	util.Encrypt.SetKey(userService.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(userService.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: userService.UserName,
		NickName: userService.NickName,
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    util.Encrypt.AesEncoding("10000"), //初始金额的加密
	}
	//密码加密
	if err = user.SetPassword(userService.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 用户登录
func (userService *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	//判断用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(userService.UserName)
	if !exist || err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请注册",
		}
	}
	//校验密码
	if user.CheckPassword(userService.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录",
		}
	}
	//token发放
	token, err := util.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "token生成失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: dto.BuildUser(user), Token: token},
	}
}

// 用户修改信息
func (service *UserService) Update(ctx context.Context, id uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	//找到这个用户
	user, err = userDao.GetUserById(id)
	//自修改nickname
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(id, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   dto.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// Post 上传头像
func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	//保存图片到本地的函数
	path, err := UpdateAvatarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	user.Avatar = path
	userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   dto.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// 发送邮箱
func (sendService SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice //模板通知
	//签发邮箱token
	emailToken, err := util.GenerateEmailToken(uId, sendService.OperationType, sendService.Email, sendService.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(sendService.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	address = conf.ValidEmail + emailToken //发送方
	emailStr := notice.Text
	emailTex := strings.Replace(emailStr, "Email", address, -1)
	message := mail.NewMessage()
	message.SetHeader("From", conf.SmtpEmail)
	message.SetHeader("To", sendService.Email)
	message.SetHeader("Subject", "FanOne")
	message.SetBody("text/html", emailTex)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(message); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 验证邮箱
func (validService ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.Success
	// 验证token
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeOut
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if operationType == 1 {
		//绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		//解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		//修改密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   dto.BuildUser(user),
	}
}

// 展示用户金额
func (showMoneyService ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   dto.BuildMoney(user, showMoneyService.Key),
	}
}
