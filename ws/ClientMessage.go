package ws

type ClientMessage struct {
	// 消息类型 1 单聊
	Type         int    `json:"type,omitempty"`
	TargetUserId uint   `json:"targetUserId,omitempty"`
	Content      string `json:"content,omitempty"`
}
