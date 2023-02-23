package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
	"im/config"
)

var (
	DB         *gorm.DB
	Config     config.Config
	Trans      ut.Translator
	HttpClient *resty.Client
)
