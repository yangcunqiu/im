package model

import "gorm.io/gorm"

const (
	Unused = iota
	Temp
	SendFail
	NoPass
	OffLine
	Pass
)

type FriendRequest struct {
	gorm.Model
	SenderUserId uint   `gorm:"type:int;not null;comment:发送者id"`
	TargetUserId uint   `gorm:"type:int;not null;comment:被添加好友的id"`
	Note         string `gorm:"type:varchar(200);comment:添加备注"`
	Status       int    `gorm:"type:int;not null;comment:状态: 1 暂存, 2 发送请求失败 3 未同意, 4 离线, 5 同意"`
}

func (f FriendRequest) TableName() string {
	return "friend_request"
}
