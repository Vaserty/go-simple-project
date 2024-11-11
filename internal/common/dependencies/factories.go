package dependencies

import (
	"go-simple-project/internal/common/config"

	logger "github.com/sirupsen/logrus"
)

func DependencyFactory(config *config.Config) *Dependency {
	dependencies := &Dependency{Config: config}
	logger.Debug("Dependecies have been successfully created.")
	return dependencies
}
