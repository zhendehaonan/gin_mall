package e

// 状态码
const (
	Success       = 200
	Error         = 500
	InvalidParams = 400 //参数错误

	//user模块的错误
	ErrorExistUser             = 30001 //用户名已存在
	ErrorFailEncryption        = 30002 //密码加密失败
	ErrorExistUserNotFound     = 30003 //用户不存在
	ErrorNotCompare            = 30004 //密码错误
	ErrorAuthToken             = 30005 //token认证失败
	ErrorAuthCheckTokenTimeOut = 30006 //token过期
	ErrorUploadFail            = 30007 //图片上传失败
	ErrorSendEmail             = 30008 //邮件发送失败

	//product模块的错误
	ErrorProductImgUpload = 40001 //商品图片上传失败
	ErrorProductCreate    = 40002 //商品创建失败

	//favorite模块的错误
	ErrorFavoriteExist = 50001 //商品已被收藏

	//address模块的错误
	ErrorAddressExist    = 60001 //地址已存在
	ErrorAddressNotExist = 60002 //地址不存在
)
