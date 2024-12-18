package xerr

// 厂商接口返回
var (
	AuthError    = NewError(20240001, "用户不属于该系统")
	GameNotFound = NewError(20240002, "游戏不存在")
)

// 平台接口返回
var (
	SignError          = NewError(20241000, "算签错误")
	UserNotExistsError = NewError(20241001, "用户不存在")
	SystemError        = NewError(20241002, "系统错误")
)

// 用户操作错误码
var (
	InvalidUserError   = NewError(20242000, "无效的数据")
	GameNotExistsError = NewError(20242001, "游戏不存在")
)
