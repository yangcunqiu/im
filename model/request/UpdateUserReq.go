package request

type UpdateUserReq struct {
	ID         uint   `json:"id"`
	Name       string `json:"name,omitempty" binding:"required,gte=5" label:"用户名"`
	ProvinceId uint   `json:"provinceId,omitempty"`
	CityId     uint   `json:"cityId,omitempty"`
	DistrictId uint   `json:"districtId,omitempty"`
}
