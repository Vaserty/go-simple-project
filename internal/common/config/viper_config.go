package config

import (
	"go-simple-project/internal/common/customtypes"

	"github.com/spf13/viper"
)

const (
	SettingsFolderPath customtypes.Path = customtypes.Path("settings")
)

func viperSetConfig(environment customtypes.Environment) {
	viper.AddConfigPath(string(SettingsFolderPath))
	viper.SetConfigType("yaml")
	viper.SetConfigName(string(environment))
}

func viperReadConfig(environment customtypes.Environment) (*Config, error) {
	viperSetConfig(environment)

	config := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return config, err
	}

	return config, nil
}
