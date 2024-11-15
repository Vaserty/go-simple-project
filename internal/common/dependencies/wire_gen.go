// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependencies

import (
	"go-simple-project/internal/common/config"
)

// Injectors from wire.go:

func InitializeDependencies() (*Dependency, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	dependency := DependencyFactory(configConfig)
	return dependency, nil
}
