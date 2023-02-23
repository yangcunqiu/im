package model

import (
	"gorm.io/gorm"
	"time"
)

type CallLog struct {
	gorm.Model
	Type          int       `gorm:"type:int;not null;comment:url方向类型, 1: im -> third, 2: third -> im"`
	Url           string    `gorm:"type:varchar(255);not null;comment:url"`
	MethodType    string    `gorm:"type:varchar(10);not null;comment:方法类型"`
	ServiceType   string    `gorm:"type:varchar(20);not null;comment:业务类型"`
	InvokeStatus  int       `gorm:"type:int;not null;comment:调用状态"`
	ServiceStatus int       `gorm:"type:int;not null;comment:业务状态"`
	RequestStr    string    `gorm:"type:longtext;comment:请求报文"`
	RequestTime   time.Time `gorm:"comment:请求时间"`
	ResponseStr   string    `gorm:"type:longtext;comment:响应报文"`
	ResponseTime  time.Time `gorm:"comment:响应时间"`
	ErrorString   string    `gorm:"type:longtext;comment:错误原因"`
	Cost          int64     `gorm:"type:int;comment:耗时(ms)"`
}
