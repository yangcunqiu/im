package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"im/handler"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(c *gin.Context) {
	// (协议升级) 升级http GET请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		handler.Fail(c, handler.WebSocketConnectionError, "")
	}
	defer ws.Close()
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

func sendAloneChat() {

}

func sendMessage(targetCoon *websocket.Conn, message Message) {
	err := targetCoon.WriteJSON(message)
	if err != nil {
		log.Println("消息发送失败, err:", err)
	}
}
