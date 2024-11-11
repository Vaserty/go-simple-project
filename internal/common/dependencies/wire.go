//go:build wireinject
// +build wireinject

package dependencies

import (
	"go-simple-project/internal/common/config"

	"github.com/google/wire"
)

func InitializeDependencies() (*Dependency, error) {
	wire.Build(
		config.NewConfig,
		DependencyFactory,
	)
	return &Dependency{}, nil
}
