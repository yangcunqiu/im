package handler

import "im/model"

var (
	// 通用错误 100
	InternalServerError      = model.ErrorOf(10000, "服务器异常")
	ParamsBindingError       = model.ErrorOf(10001, "参数填写错误")
	TokenEmptyError          = model.ErrorOf(10002, "token为空")
	TokenMalformedError      = model.ErrorOf(10003, "token错误")
	TokenExpiredError        = model.ErrorOf(10004, "token过期")
	TokenNoTActiveError      = model.ErrorOf(10005, "token是非活动的")
	TokenParseError          = model.ErrorOf(10006, "token解析失败")
	ParamsPhoneEmptyError    = model.ErrorOf(10007, "手机号不能为空")
	ParamsEmailEmptyError    = model.ErrorOf(10008, "邮箱不能为空")
	VerifyCodeError          = model.ErrorOf(10009, "验证码错误")
	OAuthVerifyError         = model.ErrorOf(10010, "权限不足")
	WebSocketConnectionError = model.ErrorOf(10011, "建立webSocket链接失败")
	// 用户模块错误 200
	UserNotFoundByIdError              = model.ErrorOf(20000, "用户不存在")
	SaveUserError                      = model.ErrorOf(20001, "保存用户出错")
	UserNameSameError                  = model.ErrorOf(20002, "用户名重复")
	UserPasswordError                  = model.ErrorOf(20003, "密码错误")
	UserLoginNameOrPasswordEmptyError  = model.ErrorOf(20004, "用户名或密码不能为空")
	UserLoginEmailOrPasswordEmptyError = model.ErrorOf(20005, "邮箱或密码不能为空")
	UserLoginPhoneOrCodeEmptyError     = model.ErrorOf(20006, "手机号或验证码不能为空")
	UserLoginTypeError                 = model.ErrorOf(20007, "登录方式错误")
	UserLoginNameOrPasswordVailError   = model.ErrorOf(20008, "用户名或密码错误")
	UserLoginError                     = model.ErrorOf(20009, "登录失败")
	UserPhoneSameError                 = model.ErrorOf(20010, "手机号已绑定")
	UserEmailSameError                 = model.ErrorOf(20011, "邮箱已绑定")
	PhoneVerifyError                   = model.ErrorOf(20012, "手机号格式错误")
	EmailVerifyError                   = model.ErrorOf(20013, "邮箱格式错误")
	UserNameEmptyError                 = model.ErrorOf(20014, "用户名不能为空")
	UserNotFoundByNameError            = model.ErrorOf(20015, "用户未找到")
	AddFriendSameError                 = model.ErrorOf(20016, "已经是好友了, 无需再次添加")
	AddFriendRequestNotFoundError      = model.ErrorOf(20017, "该好友没有向你发送过添加好友请求")
)
