package config

import (
	logger "github.com/sirupsen/logrus"
)

func NewConfig() (*Config, error) {
	environment := getEnvironment()

	config, err := viperReadConfig(environment)

	if err != nil {
		return config, err
	}

	if err := validateConfigContent(config); err != nil {
		return config, err
	}
	logger.Debug("Config has been successfully loaded.")

	configLogger(config)

	logger.Debug("Logger has been successfully configured.")

	return config, nil
}
