package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
	"im/config"
	"im/model"
)

var (
	DB         *gorm.DB
	Config     config.Config
	Trans      ut.Translator
	User       model.User
	HttpClient *resty.Client
	RDB        *redis.Client
)
