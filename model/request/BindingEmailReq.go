package request

type BindingEmailReq struct {
	UserId uint   `json:"userId" binding:"required"`
	Email  string `json:"email,omitempty" binding:"required,email"`
	Code   string `json:"code,omitempty" binding:"required,len=6"`
}
