package health

import (
	"fiber-template/cmd/constants"
	"fiber-template/internal/config"
	"fiber-template/internal/models"
	"time"
)

type HealthService interface {
	GetSimpleCheck(chanHealth chan models.ComponentCheck)
	GetDatabaseCheck(chanHealth chan models.ComponentCheck)
}

type ServiceHealt struct {
	healthRepository HealthDatabaseRepository
	config           *config.Config
}

func NewHealthService(c *config.Config, healthRepository HealthDatabaseRepository) HealthService {
	return &ServiceHealt{
		healthRepository: healthRepository,
		config:           c,
	}
}

func (*ServiceHealt) GetSimpleCheck(chanHealth chan models.ComponentCheck) {
	now := time.Now()
	details := models.NewComponentCheck("simple", "simple-check", now)
	updateHealthDetails(&details, now, nil)
	chanHealth <- details
}

func updateHealthDetails(details *models.ComponentCheck, now time.Time, err error) {
	if err == nil {
		details.Status = constants.PassHealthCheck
	} else {
		details.Status = constants.FailHealthCheck
		details.Output = err.Error()
	}
	details.ObservedValue = time.Since(now).Seconds() * 1000
	details.ObservedUnit = constants.MetricUnitHealthCheck
}

func (h *ServiceHealt) GetDatabaseCheck(chanHealth chan models.ComponentCheck) {
	now := time.Now()
	details := models.NewComponentCheck("database", "db-check", now)
	err := h.healthRepository.GetDatabaseHealth()
	updateHealthDetails(&details, now, err)
	chanHealth <- details
}
