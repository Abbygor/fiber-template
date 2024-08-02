package httpserver

import (
	"fiber-template/internal/container"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Server struct {
	Fiber        *fiber.App
	dependencies *container.Dependencies
	logger       *zerolog.Logger
}

func NewServer(depen *container.Dependencies, log *zerolog.Logger) *Server {
	return &Server{
		Fiber:        fiber.New(),
		dependencies: depen,
		logger:       log,
	}
}

func (s *Server) Start() {
	s.Fiber.Use(func(c *fiber.Ctx) error {
		s.logger.Info().Str("method", c.Method()).Str("path", c.Path()).Str("ip", c.IP()).Msg("Request received")
		return c.Next()
	})

	s.logger.Fatal().Err(s.Fiber.Listen(fmt.Sprintf(":%s", s.dependencies.Config.ProjectInfo.Port))).Msg("Server shutting down")
}

func (s *Server) SetErrorHandler() {
	s.Fiber.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error occurred")
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return nil
	})
}
