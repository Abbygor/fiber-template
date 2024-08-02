package health

import (
	"errors"
	"fiber-template/internal/config"

	"gorm.io/gorm"
)

type HealthDatabaseRepository interface {
	GetDatabaseHealth() error
}

type RepositoryHealthDatabase struct {
	config *config.Config
	db     *gorm.DB
}

func NewHealthDatabaseRepository(c *config.Config, db *gorm.DB) HealthDatabaseRepository {
	return &RepositoryHealthDatabase{
		config: c,
		db:     db,
	}
}

func (r *RepositoryHealthDatabase) GetDatabaseHealth() error {
	if r.db == nil {
		return errors.New("database connection fail")
	}
	if pinger, ok := r.db.ConnPool.(interface{ Ping() error }); ok {
		return pinger.Ping()
	}
	return nil
}
