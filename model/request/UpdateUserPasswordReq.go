package request

type UpdateUserPasswordReq struct {
	ID                 uint   `json:"id"`
	Password           string `json:"password,omitempty" binding:"required,gte=8" label:"原始密码"`
	NewPassword        string `json:"newPassword,omitempty" binding:"required,gte=8" label:"新密码"`
	ConfirmNewPassword string `json:"confirmNewPassword,omitempty" binding:"required,eqfield=NewPassword" label:"确认新密码"`
}
