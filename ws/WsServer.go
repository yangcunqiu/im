package ws

import "github.com/gorilla/websocket"

type WsServer struct {
	// 用户列表
	userCoonMap map[uint]*websocket.Conn
	coonUserMap map[*websocket.Conn]uint
	login       chan *UserCoon
	logout      chan uint
	// 广播消息 用来给所有用户发消息
	broadcast chan []byte
}

func createWsServer() *WsServer {
	return &WsServer{
		userCoonMap: make(map[uint]*websocket.Conn),
		coonUserMap: make(map[*websocket.Conn]uint),
		login:       make(chan *UserCoon),
		logout:      make(chan uint),
		broadcast:   make(chan []byte),
	}
}

func (w WsServer) getCoonByUserId(userId uint) *websocket.Conn {
	if coon, ok := w.userCoonMap[userId]; ok {
		return coon
	}
	return nil
}

func (w WsServer) run() {
	for {
		select {
		case userCoon := <-w.login:
			w.userCoonMap[userCoon.userId] = userCoon.coon
		case userId := <-w.logout:
			if _, ok := w.userCoonMap[userId]; ok {
				delete(w.userCoonMap, userId)
			}
		case data := <-w.broadcast:
			for _, coon := range w.userCoonMap {
				coon.WriteMessage(1, data)
			}
		}
	}
}
