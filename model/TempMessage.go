package model

import "time"

type TempMessage[T any] struct {
	// 消息类型 1 添加好友消息, 2 回复添加好友消息, 3 单聊, 4 群聊, 5 系统消息
	Type           int       `json:"type,omitempty"`
	SendTime       time.Time `json:"sendTime"`
	SenderUserId   uint      `json:"senderUserId,omitempty"`
	SenderUserName string    `json:"senderUserName,omitempty"`
	TargetUserId   uint      `json:"targetUserId,omitempty"`
	TargetUserName string    `json:"targetUserName,omitempty"`
	Message        T         `json:"message,omitempty"`
}
