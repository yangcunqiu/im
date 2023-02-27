package request

type AddFriend struct {
	UserId uint `json:"userId" binding:"required"`
}
