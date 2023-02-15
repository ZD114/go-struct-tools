package config

type Config struct {
	SavePath      string `yaml:"save-path"`
	MysqlUrl      string `yaml:"mysql-url"`
	EnableJsonTag bool   `yaml:"enable-json-tag"`
}

var configuration Config

func GetConfig() Config {
	return configuration
}
