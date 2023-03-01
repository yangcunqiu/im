package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"im/global"
	"im/model"
	"im/model/vo"
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

func GetFriendListByUserId(userId uint) []vo.UserFriendVO {
	f := make([]vo.UserFriendVO, 0)
	global.DB.Raw("select u.id FriendId, u.name FriendName from `user` u left join friend f on f.friend_user_id = u.id where f.user_id = ?", userId).Scan(&f)
	return f
}
