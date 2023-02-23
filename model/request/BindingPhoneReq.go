package request

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type BindingPhoneReq struct {
	UserId uint   `json:"userId" binding:"required"`
	Phone  string `json:"phone,omitempty" binding:"required,len=11,IsPhone"`
	Code   string `json:"code,omitempty" binding:"required,len=6"`
}

func IsPhone(fl validator.FieldLevel) bool {
	if phone, ok := fl.Field().Interface().(string); ok {
		re, err := regexp.Compile("^1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\\d{8}$")
		if err != nil {
			return false
		}
		return re.MatchString(phone)
	}
	return false
}
