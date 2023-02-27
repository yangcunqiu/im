package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"im/global"
	"im/handler"
	"im/model"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsServer = createWsServer()

func init() {
	go wsServer.run()
}

func WsHandler(c *gin.Context) {
	// (协议升级) 升级http GET请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		handler.Fail(c, handler.WebSocketConnectionError, "")
	}
	defer ws.Close()
	userCoon := UserCoon{
		userId: global.User.ID,
		coon:   ws,
	}
	wsServer.login <- &userCoon
	for {
		// 读取客户端数据
		mt, message, err := ws.ReadMessage()
		log.Println("--------messageType", mt)
		if err != nil {
			handler.Fail(c, handler.WebSocketConnectionError, "读取失败")
			break
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			handler.Fail(c, handler.WebSocketConnectionError, "写入失败")
			break
		}
	}
}

func AddFriend(sender *model.User, targetUserId uint) {
	addFriendReq := AddFriendReq{
		SenderId:   sender.ID,
		SenderName: sender.Name,
	}

	conn := wsServer.getCoonByUserId(targetUserId)
	if conn != nil {
		send(conn, addFriendReq)
	} else {
		// TODO 用户不在线, 暂存, 上线时推送
	}
}

func send(targetCoon *websocket.Conn, any any) {
	//err := targetCoon.WriteJSON(any)
	err := targetCoon.WriteJSON(any)
	if err != nil {
		log.Println("发送失败, err: ", err)
	}
}

func sendAloneChat() {

}

func sendMessage(targetCoon *websocket.Conn, message Message) {
	err := targetCoon.WriteJSON(message)
	if err != nil {
		log.Println("消息发送失败, err:", err)
	}
}
