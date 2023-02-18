package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"im/dao"
	"im/global"
	"im/handler"
	"im/model"
	"im/model/request"
	"im/model/vo"
	"im/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUserList(c *gin.Context) {
	handler.Success(c, dao.GetUserList())
}

func GetUser(c *gin.Context) {
	// 解析参数
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	// 查询用户
	user, err := dao.GetUser(uint(id))
	if err != nil {
		handler.Fail(c, handler.UserNotFoundByIdError)
		return
	}
	// 查询用户登录信息
	loginInfoByUserId := dao.GetUserLoginInfoByUserId(user.ID)
	userInfoVO := vo.UserInfoVO{
		ID:               user.ID,
		Name:             user.Name,
		Phone:            user.Phone,
		PhoneAttribution: user.PhoneAttribution,
		Email:            user.Email,
		IsAdmin:          user.IsAdmin,
		IsLogin:          user.IsLogin,
		ClientIP:         loginInfoByUserId.ClientIP,
		IPAttribution:    loginInfoByUserId.IPAttribution,
		LastLoginTime:    loginInfoByUserId.LastLoginTime.Time,
		LastLogoutTime:   loginInfoByUserId.LastLogoutTime.Time,
		OSVersion:        loginInfoByUserId.OSVersion,
		Browser:          loginInfoByUserId.Browser,
	}
	c.JSON(http.StatusOK, userInfoVO)
}

func RegisterUser(c *gin.Context) {
	// 绑定参数
	var createUserReq request.CreateUserReq
	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		handler.Fail(c, handler.ParamsBindingError)
		return
	}

	// TODO 校验 + 验证码 + 唯一
	// 密码加密
	// 保存用户
	// TODO 第三方接口获取其他信息

	area := model.Area{
		ProvinceId: createUserReq.ProvinceId,
		CityId:     createUserReq.CityId,
		DistrictId: createUserReq.DistrictId,
	}
	user := model.User{
		Name:     createUserReq.Name,
		Password: createUserReq.Password,
		Phone:    createUserReq.Phone,
		Email:    createUserReq.Email,
		Area:     area,
		IsAdmin:  false,
		IsLogin:  false,
	}

	// 开启事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.CreateUser(tx, &user); err != nil {
			return err
		}
		// 保存用户登录信息
		// 解析请求
		loginInfo := analysisRequest(c)
		userLoginInfo := model.UserLoginInfo{
			UserId:    user.ID,
			ClientIP:  loginInfo.ClientIP,
			OSVersion: loginInfo.OSVersion,
			Browser:   loginInfo.Browser,
		}
		if err := dao.CreateUserLoginInfo(tx, &userLoginInfo); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		handler.Fail(c, handler.SaveUserError)
		return
	}
	handler.Success(c, nil)
}

func analysisRequest(c *gin.Context) *model.UserLoginInfo {
	userAgent := c.GetHeader("user-agent")

	loginInfo := model.UserLoginInfo{
		ClientIP:  c.ClientIP(),
		OSVersion: utils.GetOSVersionByUserAgent(userAgent),
		Browser:   utils.GetBrowserByUserAgent(userAgent),
	}
	return &loginInfo
}
