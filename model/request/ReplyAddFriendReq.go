package request

type ReplyAddFriendReq struct {
	TargetUserId uint `json:"targetUserId,omitempty"`
	// 是否通过 3 未同意, 5 同意
	Status int `json:"status,omitempty"`
}
