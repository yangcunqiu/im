package ws

type AddFriendReq struct {
	SenderId      uint   `json:"senderId,omitempty"`
	SenderName    string `json:"senderName,omitempty"`
	SenderHeadUrl string `json:"senderHeadUrl,omitempty"`
	Note          string `json:"note,omitempty"`
}
