package container

import (
	"context"
	"fiber-template/internal/app/health"
	"fiber-template/internal/config"
	"fiber-template/pkg/gorm"

	"github.com/rs/zerolog"
)

type Dependencies struct {
	Config           *config.Config
	HealthController health.HealthController
}

func Build(ctx context.Context, l zerolog.Logger) Dependencies {
	dependencies := Dependencies{}

	config := config.NewConfig(ctx, l)

	dbGorm := gorm.NewDBGorm(config)
	db, err := dbGorm.Create()
	if err != nil {
		l.Fatal().Str("module", "container").Str("function", "Build").Msg(err.Error())
	}

	healthRepository := health.NewHealthDatabaseRepository(config, db)
	serviceHealth := health.NewHealthService(config, healthRepository)
	dependencies.HealthController = health.NewHealthController(serviceHealth, config)

	dependencies.Config = config

	return dependencies
}
