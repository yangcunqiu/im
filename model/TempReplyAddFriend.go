package model

type TempReplyAddFriend struct {
	ReplyUserId   uint   `json:"replyUserId,omitempty"`
	ReplyUserName string `json:"replyUserName,omitempty"`
	Status        int    `json:"status,omitempty"`
}
