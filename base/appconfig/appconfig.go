package appconfig

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`

	Database struct {
		Mysql struct {
			Url          string `yaml:"url"`
			MaxIdleConns int    `yaml:"max_idle_conns"`
			MaxOpenConns int    `yaml:"max_open_conns"`
		} `yaml:"mysql"`
	} `yaml:"database"`
}

func NewConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := new(Config)

	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
