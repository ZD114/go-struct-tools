package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"zhangda/go-tools/log"
)

type Config struct {
	SavePath      string            `yaml:"save-path"`
	MysqlUrl      string            `yaml:"mysql-url"`
	EnableJsonTag bool              `yaml:"enable-json-tag"`
	Server        ServerProperties  `yaml:"server"`
	Logging       LoggingProperties `yaml:"logging"`
	Spring        SpringProperties  `yaml:"spring"`
}

type ServerProperties struct {
	Port int32 `yaml:"port"`
}

type LoggingProperties struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type SpringProperties struct {
	Application ApplicationProperties           `yaml:"application"`
	Profiles    ProfilesProperties              `yaml:"profiles"`
	Datasource  map[string]DatasourceProperties `yaml:"datasource"`
}

type ApplicationProperties struct {
	Name string `yaml:"name"`
}

type ProfilesProperties struct {
	Active string `yaml:"active"`
}

type DatasourceProperties struct {
	Url             string `yaml:"url"`
	MaxIdleConn     int    `yaml:"maxIdleConn"`
	MaxOpenConn     int    `yaml:"maxOpenConn"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

const (
	DefaultConfigFile = "application.yml"
)

var configuration Config

func init() {
	conf := new(Config)

	f := DefaultConfigFile

	if data, err := os.ReadFile(f); err != nil {
		log.Logger.Error("config", log.Any("config", err))

		return
	} else if err = yaml.Unmarshal(data, &conf); err != nil {
		log.Logger.Error("config", log.Any("config", err))

		return
	}

	configuration = *conf
}

func GetConfig() Config {
	return configuration
}
