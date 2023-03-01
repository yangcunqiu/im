package model

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	UserId       uint   `gorm:"type:int;not null;comment:用户id"`
	FriendUserId uint   `gorm:"type:int;not null;comment:用户好友id"`
	Note         string `gorm:"type:varchar(200);comment:用户的好友备注"`
}
