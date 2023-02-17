package global

import (
	"gorm.io/gorm"
	"im/config"
)

var (
	DB     *gorm.DB
	Config config.Config
)
