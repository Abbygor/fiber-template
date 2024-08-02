package httpserver

func (s *Server) Routes() {
	s.Fiber.Get("/health", s.dependencies.HealthController.Health)
	s.Fiber.Get("/health/dependencies", s.dependencies.HealthController.HealthDependencies)
}
