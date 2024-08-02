package config

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/sethvargo/go-envconfig"
)

type (
	Config struct {
		ProjectInfo    ProjectInfo
		GormConnection GormConnection
	}
)

type ProjectInfo struct {
	Name     string `env:"PROJECT_NAME,default=fiber-template"`
	Version  string `env:"PROJECT_VERSION,default=1.0.0"`
	Language string `env:"PROJECT_LANGUAGE,default=Go"`
	Port     string `env:"PORT,default=4001"`
	Env      string `env:"ENV,default=local"`
}
type GormConnection struct {
	Server   string `env:"POSTGRES_SERVER,required"`
	Database string `env:"POSTGRES_DATABASE,required"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Port     int    `env:"POSTGRES_PORT,default=5432"`
}

func NewConfig(ctx context.Context, l zerolog.Logger) *Config {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		l.Fatal().Str("module", "envconfig").Str("function", "NewConfig").Msg(err.Error())
	}
	return &cfg
}
