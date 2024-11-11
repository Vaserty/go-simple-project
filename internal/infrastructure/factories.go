package infrastructure

import (
	"go-simple-project/internal/common/dependencies"

	logger "github.com/sirupsen/logrus"
)

func NewFileNumberRepository(deps *dependencies.Dependency) *FileNumberRepository {
	repository := &FileNumberRepository{deps: deps}
	logger.Debug("FileNumberRepository has been successfully created.")
	return repository
}
