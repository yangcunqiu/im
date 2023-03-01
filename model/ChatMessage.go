package model

import (
	"gorm.io/gorm"
	"time"
)

type ChatMessage struct {
	gorm.Model
	// 1 单聊
	Type           int       `gorm:"type:int;not null;comment:消息类型"`
	SenderUserId   uint      `gorm:"type:int;not null;comment:发送者id"`
	SenderUserName string    `gorm:"type:int;not null;comment:发送者名称"`
	SendTime       time.Time `gorm:"type:int;not null;comment:发送时间"`
	Context        string    `gorm:"type:int;not null;comment:发送内容"`
	TargetUserId   uint      `gorm:"type:int;not null;comment:目标用户id"`
}

func (c ChatMessage) TableName() string {
	return "chat_message"
}
