package ws

import "time"

type Message struct {
	sender      *UserCoon
	sendTime    time.Time
	receiver    []*UserCoon
	messageType int
	groupId     uint
	context     string
}
