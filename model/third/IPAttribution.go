package third

type IPAttribution struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Result    struct {
		Ip       string `json:"ip"`
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		AdInfo struct {
			Nation   string `json:"nation"`
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
			Adcode   int    `json:"adcode"`
		} `json:"ad_info"`
	} `json:"result"`
}
