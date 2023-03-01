package dao

import (
	"im/global"
	"im/model"
)

func AddChatMessage(msg *model.ChatMessage) {
	global.DB.Model(&model.ChatMessage{}).Create(msg)
}
