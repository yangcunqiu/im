package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"im/global"
	"im/model"
)

func AddFriend(friend model.Friend) {
	global.DB.Model(&model.Friend{}).Create(&friend)
}

func GetFriendByUserIdAndTargetUserId(userId uint, targetUserId uint) (*model.Friend, bool) {
	f := model.Friend{}
	err := global.DB.Where("user_id = ? and friend_user_id = ?", userId, targetUserId).First(&f).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false
	}
	return &f, true
}
