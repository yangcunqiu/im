package handler

import "im/model"

var (
	// 通用错误 100
	InternalServerError = model.ErrorOf(10000, "服务器异常")
	ParamsBindingError  = model.ErrorOf(10001, "参数绑定错误")
	// 用户模块错误 200
	UserNotFoundByIdError = model.ErrorOf(20000, "用户不存在")
	SaveUserError         = model.ErrorOf(20001, "保存用户出错")
)
