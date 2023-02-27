package ws

import (
	"github.com/gorilla/websocket"
)

type UserCoon struct {
	userId      uint
	userName    string
	userHeadUrl string
	coon        *websocket.Conn
}
