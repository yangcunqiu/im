package config

type Email struct {
	FromAccount              string `yaml:"fromAccount"`
	EmailServerAddr          string `yaml:"emailServerAddr"`
	Host                     string `yaml:"host"`
	Username                 string `yaml:"username"`
	AuthorizationCode        string `yaml:"authorizationCode"`
	EmailVerifyCodeExpireMin int    `yaml:"emailVerifyCodeExpireMin"`
}
