package services

import (
	"go-simple-project/internal/common/dependencies"
	"go-simple-project/internal/infrastructure"

	logger "github.com/sirupsen/logrus"
)

func newNotFoundSearchResult() *SearchResultDto {
	return &SearchResultDto{Index: -1, Value: -1, Found: false}
}

func newFoundSearchResult(index int, value int) *SearchResultDto {
	return &SearchResultDto{Index: index, Value: value, Found: true}
}

func NewSearchValueService(
	repository infrastructure.IFileNumberRepository,
	deps *dependencies.Dependency,
) (*SearchValueService, error) {
	service := &SearchValueService{deps: deps, repository: repository}
	logger.Debug("SearchValueService has been successfully created.")
	err := service.LoadValues()
	logger.Debug("Values from file have been successfully loaded to service.")
	return service, err
}
