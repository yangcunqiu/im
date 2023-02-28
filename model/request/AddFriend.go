package request

type AddFriend struct {
	UserId uint   `json:"userId" binding:"required"`
	Note   string `json:"note"`
}
