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
	"regexp"
	"strconv"
	"time"
)

var subject = "验证码"

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
		handler.Fail(c, handler.ParamsPhoneEmptyError)
	}

	// 手机号正则校验
	re, _ := regexp.Compile(global.Config.Regexp.Phone)
	if !re.MatchString(phone) {
		handler.Fail(c, handler.PhoneVerifyError)
		return
	}

	code := utils.GenerateRandomNumberToString(6)
	// 存redis
	expireMin := global.Config.Third.Aliyun.PhoneVerifyCodeExpireMin
	global.RDB.Set(context.Background(), phone, code, time.Duration(expireMin)*time.Minute)
	ok := invoke.SendVerifyCode(phone, code, expireMin)
	if !ok {
		handler.Fail(c, model.ErrorResult{}, "发送失败")
		return
	}
	handler.Success(c, "")
}

func SendEmailVerifyCode(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		handler.Fail(c, handler.ParamsEmailEmptyError)
	}

	// 邮箱正则校验
	re, _ := regexp.Compile(global.Config.Regexp.Email)
	if !re.MatchString(email) {
		handler.Fail(c, handler.EmailVerifyError)
		return
	}

	code := utils.GenerateRandomNumberToString(6)
	// 存redis
	expireMin := global.Config.Email.EmailVerifyCodeExpireMin
	global.RDB.Set(context.Background(), email, code, time.Duration(expireMin)*time.Minute)
	htmlStr := "<b>您的验证码是: " + code + ", 有效期" + strconv.Itoa(expireMin) + "分钟" + "<b>"
	if err := invoke.SendEmail(subject, htmlStr, email); err != nil {
		handler.Fail(c, model.ErrorResult{}, "发送失败")
		return
	}
	handler.Success(c, "")
}
