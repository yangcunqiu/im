package dao

import (
	"gorm.io/gorm"
	"im/global"
	"im/model"
)

func GetUserLoginInfoByUserId(userId uint) *model.UserLoginInfo {
	userLoginInfo := model.UserLoginInfo{}
	global.DB.Where("user_id", userId).First(&userLoginInfo)
	return &userLoginInfo
}

func CreateUserLoginInfo(tx *gorm.DB, userLoginInfo *model.UserLoginInfo) error {
	return tx.Create(userLoginInfo).Error
}
