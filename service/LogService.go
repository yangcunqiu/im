package service

import (
	"im/dao"
	"im/model"
)

func SaveLog(callLog model.CallLog) {
	dao.AddCallLog(callLog)
}
