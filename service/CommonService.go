package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"im/dao"
	"im/global"
	"im/handler"
	"im/invoke"
	"im/model"
	"im/utils"
	"strconv"
	"time"
)

func GetDistrictList(c *gin.Context) {
	pid, err := strconv.Atoi(c.DefaultQuery("pid", "0"))
	if err != nil {
		handler.Fail(c, handler.ParamsBindingError, "参数类型不匹配")
	}

	// 查全部
	list := dao.GetDistrictList()

	// 转成树形结构
	handler.Success(c, utils.ListToTree(list, pid))
}

func SendPhoneVerifyCode(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		handler.Fail(c, handler.ParamsPhoneEmptyError, "")
	}
	code := utils.GenerateRandomNumberToString(6)
	// 存redis
	expireMin := global.Config.Third.Aliyun.PhoneVerifyCodeExpireMin
	global.RDB.Set(context.Background(), phone, code, time.Duration(expireMin)*time.Minute)
	ok := invoke.SendVerifyCode(phone, code, expireMin)
	if ok {
		handler.Success(c, "")
	} else {
		handler.Fail(c, model.ErrorResult{}, "发送失败")
	}
}
