package config

import (
	"go-simple-project/internal/common/customtypes"
	"os"
)

func getEnvironment() customtypes.Environment {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		return customtypes.Development
	}
	return customtypes.Environment(environment)
}

func verifyFileExist(path customtypes.Path) bool {
	if _, err := os.Stat(string(path)); err != nil {
		return false
	}
	return true
}
