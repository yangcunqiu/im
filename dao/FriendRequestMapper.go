package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"im/global"
	"im/model"
)

func AddFriendRequest(req *model.FriendRequest) {
	global.DB.Model(model.FriendRequest{}).Create(req)
}

func UpdateFriendRequestStatusById(reqId uint, status int) {
	global.DB.Model(&model.FriendRequest{
		Model: gorm.Model{
			ID: reqId,
		},
	}).Update("status", status)
}

func UpdateFriendRequestNoteById(reqId uint, note string) {
	global.DB.Model(&model.FriendRequest{
		Model: gorm.Model{
			ID: reqId,
		},
	}).Update("note", note)
}

func GetFriendRequestByUserid(senderUserId uint, targetUserId uint) (req *model.FriendRequest, ok bool) {
	req = &model.FriendRequest{}
	err := global.DB.Where("sender_user_id = ? and target_user_id = ?", senderUserId, targetUserId).First(req).Error
	if err != nil {
		return req, !errors.Is(err, gorm.ErrRecordNotFound)
	}
	return req, true
}

func GetFriendRequestById(id uint) *model.FriendRequest {
	req := &model.FriendRequest{}
	global.DB.First(req, id)
	return req
}
