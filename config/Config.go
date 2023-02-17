package config

type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
}
