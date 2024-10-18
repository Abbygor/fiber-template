package container

import (
	"context"
	"fiber-template/internal/app/authors"
	"fiber-template/internal/app/books"
	"fiber-template/internal/app/health"
	"fiber-template/internal/config"
	"fiber-template/pkg/gorm"
	"fiber-template/pkg/redis"

	"github.com/rs/zerolog"
)

type Dependencies struct {
	Config            *config.Config
	HealthController  health.HealthController
	BooksController   books.BooksController
	AuthorsController authors.AuthorsController
}

func Build(ctx context.Context, l zerolog.Logger) Dependencies {
	dependencies := Dependencies{}

	config := config.NewConfig(ctx, l)

	redisClient := redis.NewRedisClient(config)
	redisConn, err := redisClient.Create(ctx)
	if err != nil {
		l.Fatal().Str("module", "container").Str("function", "Build").Msg(err.Error())
	}

	dbGorm := gorm.NewDBGorm(config)
	db, err := dbGorm.Create()
	if err != nil {
		l.Fatal().Str("module", "container").Str("function", "Build").Msg(err.Error())
	}

	healthRepository := health.NewHealthDatabaseRepository(config, db)
	serviceHealth := health.NewHealthService(config, healthRepository)
	dependencies.HealthController = health.NewHealthController(serviceHealth, config)

	booksRepository := books.NewBooksRepository(config, db, redisConn, l)
	booksService := books.NewBooksService(booksRepository, l)
	dependencies.BooksController = books.NewBooksController(booksService, l)

	authorsRepository := authors.NewAuthorsRepository(config, db, redisConn)
	authorsService := authors.NewAuthorsService(authorsRepository)
	dependencies.AuthorsController = authors.NewAuthorsController(authorsService)

	dependencies.Config = config

	return dependencies
}
