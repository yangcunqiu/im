package ws

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/copier"
	"im/dao"
	"im/global"
	"im/handler"
	"im/middlewares"
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
	// 看下redis暂存的消息 有没有当前用户的
	key := "temp-" + strconv.Itoa(int(global.User.ID))
	tempMessage, _ := global.RDB.Get(context.Background(), key).Result()
	if tempMessage != "" {
		messageList := make([]string, 0)
		json.Unmarshal([]byte(tempMessage), &messageList)
		for _, temp := range messageList {
			if temp != "" {
				send(ws, temp)
			}
		}
		// 删掉
		global.RDB.Del(context.Background(), key)
	}

	go processTempChat(ws)

	for {
		// 读取客户端数据
		mt, message, err := ws.ReadMessage()
		middlewares.ValidationToken()(c)
		log.Println("当前发送消息用户, userId: ", global.User.ID)
		if mt == -1 {
			wsServer.logout <- global.User.ID
		}
		if err != nil {
			break
		}
		cm := ClientMessage{}
		json.Unmarshal(message, &cm)
		go processClientMessage(&cm, &global.User, ws)
	}
}

func processTempChat(ws *websocket.Conn) {
	key := "tempChat-" + strconv.Itoa(int(global.User.ID))
	tempMessage, _ := global.RDB.Get(context.Background(), key).Result()
	if tempMessage != "" {
		msgList := make([]Message, 0)
		json.Unmarshal([]byte(tempMessage), &msgList)
		for _, msg := range msgList {
			send(ws, msg)
		}
		// 删掉
		global.RDB.Del(context.Background(), key)
	}

}

func processClientMessage(cm *ClientMessage, user *model.User, userConn *websocket.Conn) {
	if cm.Type == 1 {
		_, ok := dao.GetFriendByUserIdAndTargetUserId(user.ID, cm.TargetUserId)
		if !ok {
			userConn.WriteMessage(1, []byte("对方不是你的好友, 请先添加好友再聊天吧"))
			return
		}

		msg := Message{
			Type:           1,
			SenderUserId:   user.ID,
			SenderUserName: user.Name,
			SendTime:       time.Now(),
			Context:        cm.Content,
			TargetUserId:   cm.TargetUserId,
		}

		cmsg := model.ChatMessage{}
		copier.Copy(&cmsg, &msg)
		dao.AddChatMessage(&cmsg)

		conn := wsServer.getCoonByUserId(cm.TargetUserId)
		if conn != nil {
			send(conn, msg)
		} else {
			// 暂存
			messageList := make([]Message, 0)

			key := "tempChat-" + strconv.Itoa(int(cm.TargetUserId))
			tempMessage, _ := global.RDB.Get(context.Background(), key).Result()
			if tempMessage != "" {
				json.Unmarshal([]byte(tempMessage), &messageList)
				messageList = append(messageList, msg)
			} else {
				messageList = append(messageList, msg)
			}
			marshal, _ := json.Marshal(messageList)
			global.RDB.Set(context.Background(), key, string(marshal), time.Duration(-1)*time.Minute)
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
		messageList := make([]string, 0)

		key := "temp-" + strconv.Itoa(int(targetUserId))
		tempMessage, _ := global.RDB.Get(context.Background(), key).Result()
		if tempMessage != "" {
			json.Unmarshal([]byte(tempMessage), &messageList)
			messageList = append(messageList, string(bytes))
		} else {
			messageList = append(messageList, string(bytes))
		}

		marshal, _ := json.Marshal(messageList)
		global.RDB.Set(context.Background(), key, string(marshal), time.Duration(-1)*time.Minute)
	}
}

func ReplyAddFriendRequest(replyUser *model.User, targetUserId uint, status int) {
	replyAddFriend := ReplyAddFriend{
		ReplyUserId:   replyUser.ID,
		ReplyUserName: replyUser.Name,
		Status:        status,
	}

	conn := wsServer.getCoonByUserId(targetUserId)
	if conn != nil {
		send(conn, replyAddFriend)
	} else {
		m := model.TempMessage[model.TempReplyAddFriend]{
			Type:           2,
			SendTime:       time.Now(),
			SenderUserId:   replyUser.ID,
			SenderUserName: replyUser.Name,
			TargetUserId:   targetUserId,
			Message: model.TempReplyAddFriend{
				ReplyUserId:   replyUser.ID,
				ReplyUserName: replyUser.Name,
				Status:        status,
			},
		}
		bytes, _ := json.Marshal(m)
		messageList := make([]string, 0)

		key := "temp-" + strconv.Itoa(int(targetUserId))
		tempMessage, _ := global.RDB.Get(context.Background(), key).Result()
		if tempMessage != "" {
			json.Unmarshal([]byte(tempMessage), &messageList)
			messageList = append(messageList, string(bytes))
		} else {
			messageList = append(messageList, string(bytes))
		}

		marshal, _ := json.Marshal(messageList)
		global.RDB.Set(context.Background(), key, string(marshal), time.Duration(-1)*time.Minute)
	}

}

func send(targetCoon *websocket.Conn, any any) error {
	// err := targetCoon.WriteJSON(any)
	return targetCoon.WriteJSON(any)
}
