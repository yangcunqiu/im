package config

type Aliyun struct {
	AppCode                       string `yaml:"appCode"`
	PhoneAttributionUrl           string `yaml:"phoneAttributionUrl"`
	IPAttributionUrl              string `yaml:"ipAttributionUrl"`
	PhoneSendVerifyCodeUrl        string `yaml:"phoneSendVerifyCodeUrl"`
	PhoneSendVerifyCodeTemplateId string `yaml:"phoneSendVerifyCodeTemplateId"`
	PhoneVerifyCodeExpireMin      int    `yaml:"phoneVerifyCodeExpireMin"`
}
