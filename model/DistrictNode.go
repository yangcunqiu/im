package model

// 省市区
type DistrictNode struct {
	ID                    uint
	Pid                   uint
	Name                  string
	Level                 uint8
	DistrictWrapChildList []DistrictNode
}
