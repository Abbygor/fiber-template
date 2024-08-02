package health

import (
	"fiber-template/cmd/constants"
	"fiber-template/internal/config"
	"fiber-template/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthController interface {
	Health(ctx *fiber.Ctx) error
	HealthDependencies(ctx *fiber.Ctx) error
}

type ControllerHealth struct {
	healthService HealthService
	config        *config.Config
}

func NewHealthController(service HealthService, c *config.Config) HealthController {
	return &ControllerHealth{
		healthService: service,
		config:        c,
	}
}

func (h *ControllerHealth) Health(ctx *fiber.Ctx) error {
	health := models.NewHealthCheck(h.config.ProjectInfo.Version, constants.PassHealthCheck)
	chanHealth := make(chan models.ComponentCheck)
	defer close(chanHealth)
	go h.healthService.GetSimpleCheck(chanHealth)
	simpleHealth := <-chanHealth
	health.Checks["simple:here-i-am"] = simpleHealth
	status := http.StatusOK
	if simpleHealth.Status == constants.FailHealthCheck {
		status = http.StatusServiceUnavailable
	}
	return ctx.Status(status).JSON(health)
}

func (h *ControllerHealth) HealthDependencies(ctx *fiber.Ctx) error {
	health := models.NewHealthCheck(h.config.ProjectInfo.Version, constants.PassHealthCheck)
	chanIcarusHealth := make(chan models.ComponentCheck)
	chanDatabaseHealth := make(chan models.ComponentCheck)

	defer closeChannels(chanIcarusHealth, chanDatabaseHealth)
	go h.healthService.GetDatabaseCheck(chanDatabaseHealth)

	databaseHealth := <-chanDatabaseHealth
	health.Checks["database:health-check"] = databaseHealth

	status := http.StatusOK
	if databaseHealth.Status == constants.FailHealthCheck {
		status = http.StatusServiceUnavailable
		health.Status = constants.FailHealthCheck
	}

	return ctx.Status(status).JSON(health)
}

func closeChannels(channels ...chan models.ComponentCheck) {
	for _, item := range channels {
		close(item)
	}
}
