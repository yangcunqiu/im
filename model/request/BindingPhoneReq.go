package request

import (
	"github.com/go-playground/validator/v10"
	"im/global"
	"regexp"
)

type BindingPhoneReq struct {
	UserId uint   `json:"userId" binding:"required"`
	Phone  string `json:"phone,omitempty" binding:"required,len=11,IsPhone"`
	Code   string `json:"code,omitempty" binding:"required,len=6"`
}

func IsPhone(fl validator.FieldLevel) bool {
	if phone, ok := fl.Field().Interface().(string); ok {
		re, err := regexp.Compile(global.Config.Regexp.Phone)
		if err != nil {
			return false
		}
		return re.MatchString(phone)
	}
	return false
}
