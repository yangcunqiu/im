package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"im/dao"
	"im/global"
	"im/handler"
	"im/invoke"
	"im/model"
	"im/model/request"
	"im/model/vo"
	"im/utils"
	"log"
	"strconv"
)

func GetUserList(c *gin.Context) {
	// 分页
	pageNum, pageSize, offset := handler.PageParams(c)
	name := c.Query("name")

	// 查询
	userList, total := dao.GetUserList(offset, pageSize, name)

	// 转换
	userVOList := make([]vo.UserInfoVO, 0)
	for _, user := range userList {
		userVO := vo.UserInfoVO{
			ID:               user.ID,
			Name:             user.Name,
			Phone:            user.Phone,
			PhoneAttribution: user.PhoneAttribution,
			Email:            user.Email,
			IsAdmin:          user.IsAdmin,
			IsLogin:          user.IsLogin,
		}
		userVOList = append(userVOList, userVO)
	}

	handler.Success(c, handler.PageOf(pageNum, pageSize, total, userVOList))
}

// 获取用户详情
func GetUser(c *gin.Context) {
	// 解析参数
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	// 查询用户
	user, err := dao.GetUser(uint(id))
	if err != nil {
		log.Println(err)
		handler.Fail(c, handler.UserNotFoundByIdError, "")
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
	handler.Success(c, userInfoVO)
}

// 用户注册
func RegisterUser(c *gin.Context) {
	// 绑定参数
	var createUserReq request.CreateUserReq
	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		handler.Fail(c, handler.ParamsBindingError, handler.SimpleValidateErrorTran(err))
		return
	}

	_, err := dao.GetUserByName(createUserReq.Name)
	if err == nil {
		// 重复
		handler.Fail(c, handler.UserNameSameError, "")
		return
	}

	// 密码加密
	salt := utils.GetUUID()
	password := utils.EncodeBySHA256(createUserReq.Password, salt)

	// 保存用户
	// TODO 第三方接口获取其他信息

	area := model.Area{
		ProvinceId: createUserReq.ProvinceId,
		CityId:     createUserReq.CityId,
		DistrictId: createUserReq.DistrictId,
	}
	user := model.User{
		Name:     createUserReq.Name,
		Password: password,
		Salt:     salt,
		Area:     area,
		IsAdmin:  false,
		IsLogin:  false,
	}

	// 开启事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.CreateUser(tx, &user); err != nil {
			return err
		}
		// 保存用户登录信息
		// 解析请求
		loginInfo := analysisRequest(c)
		userLoginInfo := model.UserLoginInfo{
			UserId:        user.ID,
			ClientIP:      loginInfo.ClientIP,
			IPAttribution: invoke.QueryIPAttribution(loginInfo.ClientIP),
			OSVersion:     loginInfo.OSVersion,
			Browser:       loginInfo.Browser,
		}
		if err := dao.CreateUserLoginInfo(tx, &userLoginInfo); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		handler.Fail(c, handler.SaveUserError, "")
		return
	}
	handler.Success(c, nil)
}

// 解析用户请求中其他信息 ip 操作系统 浏览器
func analysisRequest(c *gin.Context) *model.UserLoginInfo {
	userAgent := c.GetHeader("user-agent")

	loginInfo := model.UserLoginInfo{
		ClientIP:  c.ClientIP(),
		OSVersion: utils.GetOSVersionByUserAgent(userAgent),
		Browser:   utils.GetBrowserByUserAgent(userAgent),
	}
	return &loginInfo
}

// 修改用户 只能修改用户名和地区
func UpdateUser(c *gin.Context) {
	var updateUserReq request.UpdateUserReq
	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		handler.Fail(c, handler.ParamsBindingError, handler.SimpleValidateErrorTran(err))
		return
	}

	_, err := dao.GetUserByName(updateUserReq.Name)
	if err == nil {
		// 重复
		handler.Fail(c, handler.UserNameSameError, "")
		return
	}

	area := model.Area{
		ProvinceId: updateUserReq.ProvinceId,
		CityId:     updateUserReq.CityId,
		DistrictId: updateUserReq.DistrictId,
	}
	user := model.User{
		Model: gorm.Model{
			ID: updateUserReq.ID,
		},
		Name: updateUserReq.Name,
		Area: area,
	}
	dao.UpdateUser(&user)
	handler.Success(c, nil)
}

// 修改密码
func UpdatePassword(c *gin.Context) {
	var updateUserPasswordReq request.UpdateUserPasswordReq
	if err := c.ShouldBindJSON(&updateUserPasswordReq); err != nil {
		handler.Fail(c, handler.ParamsBindingError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 校验密码是否一致
	userById, err := dao.GetUser(updateUserPasswordReq.ID)
	if err != nil {
		handler.Fail(c, handler.UserNotFoundByIdError, "")
		return
	}

	if ok := utils.VailPasswordBySHA256(updateUserPasswordReq.Password, userById.Password, userById.Salt); !ok {
		handler.Fail(c, handler.UserPasswordError, "")
		return
	}

	// 校验通过 修改密码
	// 密码加密
	salt := utils.GetUUID()
	password := utils.EncodeBySHA256(updateUserPasswordReq.NewPassword, salt)
	user := model.User{
		Model: gorm.Model{
			ID: userById.ID,
		},
		Password: password,
		Salt:     salt,
	}
	dao.UpdateUser(&user)

	// TODO 修改成功 手机号发短信 邮箱发短信
	handler.Success(c, "")
}

func Login(c *gin.Context) {
	var userLoginReq request.UserLoginReq
	if err := c.ShouldBindJSON(&userLoginReq); err != nil {
		handler.Fail(c, handler.ParamsBindingError, "")
		return
	}

	// 根据登录方式校验必填
	if !vailLoginType(c, userLoginReq) {
		return
	}

	// TODO 多次登录失败错误 锁定账号 + 每次登录更新ip和归属 以及loginInfo中的最后登录时间
	loginType := userLoginReq.LoginType
	errorRes := new(model.ErrorResult)
	user := &model.User{}
	if loginType == 1 {
		// 用户名 密码登录
		user, errorRes = loginByNameAndPassword(userLoginReq)
	}

	if errorRes != nil {
		handler.Fail(c, *errorRes, "")
		return
	}
	// 生成token
	fmt.Println(user)
}

func loginByNameAndPassword(req request.UserLoginReq) (*model.User, *model.ErrorResult) {
	user := &model.User{}
	user, err := dao.GetUserByName(req.Name)
	if err != nil {
		return user, &handler.UserLoginNameOrPasswordVailError
	}
	encodePassword := utils.EncodeBySHA256(req.Password, user.Salt)
	if encodePassword != user.Password {
		return user, &handler.UserLoginNameOrPasswordVailError
	}
	return user, nil
}

func vailLoginType(c *gin.Context, userLoginReq request.UserLoginReq) bool {
	name := userLoginReq.Name
	email := userLoginReq.Email
	password := userLoginReq.Password
	phone := userLoginReq.Phone
	code := userLoginReq.Code
	switch userLoginReq.LoginType {
	case 1:
		if name == "" || password == "" {
			handler.Fail(c, handler.UserLoginNameOrPasswordEmptyError, "")
			return false
		}
	case 2:
		if email == "" || password == "" {
			handler.Fail(c, handler.UserLoginEmailOrPasswordEmptyError, "")
			return false
		}
	case 3:
		if phone == "" || code == "" {
			handler.Fail(c, handler.UserLoginPhoneOrCodeEmptyError, "")
			return false
		}
	default:
		handler.Fail(c, handler.UserLoginTypeError, "")
		return false
	}
	return true
}

func BindingPhone(c *gin.Context) {
	var bindingPhoneReq request.BindingPhoneReq
	if err := c.ShouldBindJSON(&bindingPhoneReq); err != nil {
		handler.Fail(c, handler.ParamsBindingError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 手机号重复校验
	handler.Success(c, bindingPhoneReq.Phone)
}
