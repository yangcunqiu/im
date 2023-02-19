package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name             string `gorm:"type:varchar(255);not null;comment:用户名"`
	Password         string `gorm:"type:varchar(256);not null;comment:密码"`
	Salt             string `gorm:"type:char(32);not null;comment:salt"`
	Phone            string `gorm:"type:varchar(255);comment:手机号"`
	PhoneAttribution string `gorm:"type:varchar(255);comment:手机号归属地"`
	Email            string `gorm:"type:varchar(255);comment:邮箱"`
	Area
	IsAdmin bool `gorm:"comment:是否是管理员"`
	IsLogin bool `gorm:"comment:当前是否登录"`
}

func (user *User) TableName() string {
	return "user"
}
