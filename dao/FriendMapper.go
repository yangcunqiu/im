package dao

import (
	"im/global"
	"im/model"
)

func AddFriend(friend model.Friend) {
	global.DB.Model(&model.Friend{}).Create(&friend)
}
