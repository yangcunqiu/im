package ws

import "time"

type Message struct {
	// 1 单聊
	Type           int       `json:"type,omitempty"`
	SenderUserId   uint      `json:"senderUserId,omitempty"`
	SenderUserName string    `json:"senderUserName,omitempty"`
	SendTime       time.Time `json:"sendTime"`
	Context        string    `json:"context,omitempty"`
	TargetUserId   uint      `json:"targetUserId,omitempty"`
}
