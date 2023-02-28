package ws

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"im/dao"
	"im/global"
	"im/handler"
	"im/model"
	"log"
	"net/http"
	"strconv"
	"time"
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

func AddFriend(sender *model.User, targetUserId uint, friendRequestId uint) {
	friendRequest := dao.GetFriendRequestById(friendRequestId)
	addFriendReq := AddFriendReq{
		SenderId:   sender.ID,
		SenderName: sender.Name,
		Note:       friendRequest.Note,
	}

	conn := wsServer.getCoonByUserId(targetUserId)
	if conn != nil {
		if err := send(conn, addFriendReq); err != nil {
			dao.UpdateFriendRequestStatusById(friendRequestId, model.SendFail)
		}
	} else {
		// 用户不在线, 上线时推送
		dao.UpdateFriendRequestStatusById(friendRequestId, model.OffLine)

		key := "mm/temp-" + strconv.Itoa(int(targetUserId))
		m := model.TempMessage[model.TempAddFriend]{
			Type:           1,
			SendTime:       time.Now(),
			SenderUserId:   sender.ID,
			SenderUserName: sender.Name,
			TargetUserId:   targetUserId,
			Message: model.TempAddFriend{
				Note: friendRequest.Note,
			},
		}
		bytes, _ := json.Marshal(m)
		global.RDB.Set(context.Background(), key, string(bytes), time.Duration(-1)*time.Minute)
	}
}

func send(targetCoon *websocket.Conn, any any) error {
	//err := targetCoon.WriteJSON(any)
	return targetCoon.WriteJSON(any)
}

func sendAloneChat() {

}

func sendMessage(targetCoon *websocket.Conn, message Message) {
	err := targetCoon.WriteJSON(message)
	if err != nil {
		log.Println("消息发送失败, err:", err)
	}
}
