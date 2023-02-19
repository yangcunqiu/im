package vo

import (
	"time"
)

type UserInfoVO struct {
	ID               uint      `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Phone            string    `json:"phone,omitempty"`
	PhoneAttribution string    `json:"phoneAttribution,omitempty"`
	Email            string    `json:"email,omitempty"`
	IsAdmin          bool      `json:"isAdmin,omitempty"`
	IsLogin          bool      `json:"isLogin,omitempty"`
	ProvinceId       uint      `gorm:"comment:省id"`
	ProvinceName     uint      `gorm:"comment:省名称"`
	CityId           uint      `gorm:"comment:市id"`
	CityName         uint      `gorm:"comment:市名称"`
	DistrictId       uint      `gorm:"comment:区id"`
	DistrictName     uint      `gorm:"comment:区名称"`
	ClientIP         string    `json:"clientIP,omitempty"`
	IPAttribution    string    `json:"IPAttribution,omitempty"`
	LastLoginTime    time.Time `json:"lastLoginTime"`
	LastLogoutTime   time.Time `json:"lastLogoutTime"`
	OSVersion        string    `json:"OSVersion,omitempty"`
	Browser          string    `json:"browser,omitempty"`
}
