package dao

import (
	"im/global"
	"im/model"
)

func AddCallLog(callLog model.CallLog) {
	global.DB.Model(model.CallLog{}).Create(&callLog)
}
