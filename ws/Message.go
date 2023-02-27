package ws

import "time"

type Message struct {
	// 1 加好友, 2 单聊, 3 群聊
	messageType int
	sender      *UserCoon
	sendTime    time.Time
	receiver    []*UserCoon
	groupId     uint
	context     string
}
