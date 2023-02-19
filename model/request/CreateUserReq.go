package request

type CreateUserReq struct {
	ID              uint   `json:"id"`
	Name            string `json:"name,omitempty" binding:"required,gte=5" label:"用户名"`
	Password        string `json:"password,omitempty" binding:"required,gte=8" label:"密码"`
	ConfirmPassword string `json:"confirmPassword,omitempty" binding:"required,eqfield=Password" label:"确认密码"`
	ProvinceId      uint   `json:"provinceId,omitempty"`
	CityId          uint   `json:"cityId,omitempty"`
	DistrictId      uint   `json:"districtId,omitempty"`
}
