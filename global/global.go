package global

import (
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
	"im/config"
	"im/model"
)

var (
	DB     *gorm.DB
	Config config.Config
	Trans  ut.Translator
	User   model.User
)
