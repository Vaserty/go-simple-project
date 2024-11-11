package config

import (
	logger "github.com/sirupsen/logrus"
)

func configLogger(config *Config) {
	level, _ := logger.ParseLevel(config.Logger.Level)
	logger.SetLevel(level)
}
