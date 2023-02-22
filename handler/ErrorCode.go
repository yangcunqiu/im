package handler

import "im/model"

var (
	// 通用错误 100
	InternalServerError = model.ErrorOf(10000, "服务器异常")
	ParamsBindingError  = model.ErrorOf(10001, "参数填写错误")
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
)
