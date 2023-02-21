package dao

import (
	"im/global"
	"im/model"
)

func GetDistrictList() []model.TreeNode {
	districtList := make([]model.TreeNode, 0)
	global.DB.Raw("select * from district").Scan(&districtList)
	return districtList
}
