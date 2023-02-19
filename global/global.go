package global

import (
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
	"im/config"
)

var (
	DB     *gorm.DB
	Config config.Config
	Trans  ut.Translator
)
