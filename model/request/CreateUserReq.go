package request

type CreateUserReq struct {
	Name       string `json:"name,omitempty"`
	Password   string `json:"password,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Code       string `json:"code"`
	Email      string `json:"email,omitempty"`
	ProvinceId uint   `json:"provinceId,omitempty"`
	CityId     uint   `json:"cityId,omitempty"`
	DistrictId uint   `json:"districtId,omitempty"`
}
