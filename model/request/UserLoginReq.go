package request

type UserLoginReq struct {
	// 登录方式 (1: 用户名密码, 2: 邮箱密码, 3:手机号验证码)
	LoginType uint8  `json:"loginType"`
	Name      string `json:"name,omitempty"`
	Password  string `json:"password,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Code      string `json:"code,omitempty"`
	Email     string `json:"email,omitempty"`
}
