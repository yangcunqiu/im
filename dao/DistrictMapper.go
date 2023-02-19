package dao

import (
	"im/global"
	"im/model"
)

func GetDistrictList() []model.District {
	districtList := make([]model.District, 0)
	global.DB.Find(&districtList)
	return districtList
}
