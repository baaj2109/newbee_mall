package address_request

type AddAddressParam struct {
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	DefaultFlag   byte   `json:"defaultFlag"` // 0-不是 1-是
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
}
