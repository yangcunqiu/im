package request

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type BindingPhoneReq struct {
	Phone string `json:"phone,omitempty" binding:"required,gte=8,IsPhone"`
	Code  string `json:"code,omitempty"`
}

func IsPhone(fl validator.FieldLevel) bool {
	if phone, ok := fl.Field().Interface().(string); ok {
		re, err := regexp.Compile("/^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$/")
		if err != nil {
			return false
		}
		return re.MatchString(phone)
	}
	return false
}
