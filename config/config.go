package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Service struct {
	Port   string `mapstructure:"port"`
	Debug  bool   `mapstructure:"debug"`
	ZooApi string `mapstructure:"zooApi"`
}

func (s *Service) Ok() {
	fmt.Print(s)
}

func LoadFromFile(configPath string) (*Service, error) {
	viper.SetConfigType("json")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config %s. error: %w", configPath, err)
	}

	config := new(Service)
	err = viper.UnmarshalExact(&config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
