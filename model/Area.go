package model

type Area struct {
	ProvinceId uint `gorm:"comment:省id"`
	CityId     uint `gorm:"comment:市id"`
	DistrictId uint `gorm:"comment:区id"`
}
