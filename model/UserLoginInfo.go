package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type UserLoginInfo struct {
	gorm.Model
	UserId         uint         `gorm:"not null;comment:user.id"`
	ClientIP       string       `gorm:"type:varchar(12);comment:客户端ip"`
	IPAttribution  string       `gorm:"type:varchar(50);comment:客户端ip归属地"`
	LastLoginTime  sql.NullTime `gorm:"comment:最后一次登录时间"`
	LastLogoutTime sql.NullTime `gorm:"comment:最后一次登出时间"`
	HeartbeatTime  sql.NullTime `gorm:"comment:心跳时间"`
	OSVersion      string       `gorm:"type:varchar(50);comment:客户端登录操作系统版本"`
	Browser        string       `gorm:"type:varchar(50);comment:客户端登录浏览器"`
}

func (userLoginInfo UserLoginInfo) TableName() string {
	return "user_login_info"
}
