package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"zhangda/go-tools/log"
)

const (
	DefaultConfigFile = "application.yml"
)

type SpringConfiguration struct {
	Spring         SpringProperties  `yaml:"spring"`
	DateFormat     string            `yaml:"date-format"`
	DateTimeFormat string            `yaml:"date-time-format"`
	TimeZone       string            `yaml:"time-zone"`
	Server         ServerProperties  `yaml:"server"`
	Logging        LoggingProperties `yaml:"logging"`
}

type SpringProperties struct {
	Application ApplicationProperties           `yaml:"application"`
	Profiles    ProfilesProperties              `yaml:"profiles"`
	DataSource  map[string]DatasourceProperties `yaml:"datasource"`
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

type ServerProperties struct {
	Port int32 `yaml:"port"`
}

type LoggingProperties struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

var springConfiguration SpringConfiguration

func init() {
	conf := new(SpringConfiguration)

	f := DefaultConfigFile

	if data, err := os.ReadFile(f); err != nil {
		log.Logger.Error("config", log.Any("config", err))

		return
	} else if err = yaml.Unmarshal(data, &conf); err != nil {
		log.Logger.Error("config", log.Any("config", err))

		return
	}

	springConfiguration = *conf
}

func GetConfig() SpringConfiguration {
	return springConfiguration
}
