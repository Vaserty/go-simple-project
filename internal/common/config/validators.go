package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func validateConfigContent(config *Config) error {
	validate := validator.New()

	if err := validate.Struct(config); err != nil {
		return err
	}

	if !verifyFileExist(config.NumbersFile.SourcePath) {
		return fmt.Errorf(
			"file with given path '%s' does not exist",
			config.NumbersFile.SourcePath,
		)
	}

	return nil
}
