package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"im/global"
	"strings"
)

func GetJoinStringByValidateErrorTran(validateErrorTran validator.ValidationErrorsTranslations) string {
	str := ""
	for _, v := range validateErrorTran {
		str += fmt.Sprintf("%v, ", v)
	}
	return strings.TrimRight(str, ", ")
}

func SimpleValidateErrorTran(err error) string {
	if err != nil {
		// 转换成validate校验错误
		errs := err.(validator.ValidationErrors)
		// 使用翻译器翻译错误
		translate := errs.Translate(global.Trans)
		return GetJoinStringByValidateErrorTran(translate)
	}
	return err.Error()
}
