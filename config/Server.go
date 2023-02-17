package config

type Server struct {
	Port        uint   `yaml:"port"`
	ContextPath string `yaml:"contextPath"`
}
