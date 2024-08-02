package gorm

import (
	"fiber-template/internal/config"
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	envProd = "prod"
)

type DBGorm interface {
	Create() (*gorm.DB, error)
}

type dbGorm struct {
	config *config.Config
}

func NewDBGorm(cfg *config.Config) DBGorm {
	return &dbGorm{
		config: cfg,
	}
}

func (d *dbGorm) Create() (*gorm.DB, error) {
	connectionString := d.buildConnectionString()

	db, err := d.initializeDBSession(connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *dbGorm) buildConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		d.config.GormConnection.Server,
		d.config.GormConnection.User,
		url.QueryEscape(d.config.GormConnection.Password),
		d.config.GormConnection.Database,
		d.config.GormConnection.Port,
	)
}

func (d *dbGorm) initializeDBSession(connectionString string) (*gorm.DB, error) {
	//var logMode = logger.Error
	if d.config.ProjectInfo.Env == envProd {
		//logMode = logger.Silent
	}
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true, // skip the snake_casing of names
		},
		//Logger: logger.Default.LogMode(logMode)
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
