package config

import "go-simple-project/internal/common/customtypes"

type Config struct {
	App struct {
		CurrentEnvironment customtypes.Environment `validate:"required,oneof=development test"`
		Name               string                  `validate:"required"`
	}
	NumbersFile struct {
		SourcePath customtypes.Path `validate:"required"`
	}
	Logger struct {
		Level string `validate:"required,oneof=debug info error"`
	}
	Searcher struct {
		MaxToleranceDiffPercent float64 `validate:"required,gte=0,lte=100"`
	}
	HttpServer struct {
		Port int `validate:"required"`
	}
}
