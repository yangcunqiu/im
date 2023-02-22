package config

type Jwt struct {
	Secret         string `yaml:"secret"`
	ExpirationTime int64  `yaml:"expirationTime"`
}
