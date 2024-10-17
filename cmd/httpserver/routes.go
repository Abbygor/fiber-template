package httpserver

func (s *Server) Routes() {
	s.Fiber.Get("/health", s.dependencies.HealthController.Health)
	s.Fiber.Get("/health/dependencies", s.dependencies.HealthController.HealthDependencies)

	s.Fiber.Post("/books", s.dependencies.BooksController.CreateBook)
	s.Fiber.Get("/books/:id", s.dependencies.BooksController.GetBookByID)
	s.Fiber.Get("/books/author/:id", s.dependencies.BooksController.GetBooksByAuthorID)
	s.Fiber.Get("/books", s.dependencies.BooksController.GetBooks)
	s.Fiber.Put("/books/:id", s.dependencies.BooksController.UpdateBook)
	s.Fiber.Delete("/books/:id", s.dependencies.BooksController.DeleteBook)
}
