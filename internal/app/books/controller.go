package books

import (
	"fiber-template/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type BooksController interface {
	CreateBook(ctx *fiber.Ctx) error
	GetBookByID(ctx *fiber.Ctx) error
	GetBooksByAuthorID(ctx *fiber.Ctx) error
	GetBooks(ctx *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
}

type ControllerBooks struct {
	booksService BooksService
	l            zerolog.Logger
}

func NewBooksController(booksService BooksService, log zerolog.Logger) BooksController {
	return &ControllerBooks{
		booksService: booksService,
		l:            log,
	}
}

func (c *ControllerBooks) CreateBook(ctx *fiber.Ctx) error {
	var book models.Book
	if err := ctx.BodyParser(&book); err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "CreateBook").Msg("error parsing book")
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdBook, err := c.booksService.CreateBook(book)
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "CreateBook").Msg("error creating book from service")
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(createdBook)
}

func (c *ControllerBooks) GetBookByID(ctx *fiber.Ctx) error {
	bookIDStr := ctx.Params("id")
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "GetBookByID").Str("book_id", bookIDStr).Msg("error parsing bookID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	book, err := c.booksService.GetBookByID(ctx.Context(), int(bookID))
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "GetBookByID").Int("book_id", int(bookID)).Msg("error getting book from service")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(book) // Retorna el libro encontrado
}

func (c *ControllerBooks) GetBooksByAuthorID(ctx *fiber.Ctx) error {
	authorIDStr := ctx.Params("id")
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "GetBooksByAuthorID").Str("book_id", authorIDStr).Msg("error parsing authorID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	books, err := c.booksService.GetBooksByAuthorID(int(authorID))
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "GetBooksByAuthorID").Str("book_id", authorIDStr).Msg("error getting books by author from service")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(books)
}
func (c *ControllerBooks) GetBooks(ctx *fiber.Ctx) error {
	books, err := c.booksService.GetBooks()
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "GetBooks").Msg("error getting all books from service")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(books)
}

func (c *ControllerBooks) UpdateBook(ctx *fiber.Ctx) error {
	bookIDStr := ctx.Params("id")
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "UpdateBook").Str("book_id", bookIDStr).Msg("error parsing bookID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var book models.Book
	if err := ctx.BodyParser(&book); err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "UpdateBook").Msg("error parsing book")
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	updatedBook, err := c.booksService.UpdateBook(int(bookID), &book)
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "UpdateBook").Int("book_id", int(bookID)).Msg("error updating book from service")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(updatedBook)
}

func (c *ControllerBooks) DeleteBook(ctx *fiber.Ctx) error {
	bookIDStr := ctx.Params("id")
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "DeleteBook").Str("book_id", bookIDStr).Msg("error parsing bookID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = c.booksService.DeleteBook(int(bookID))
	responseStatus := fiber.StatusInternalServerError
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			responseStatus = fiber.StatusNotFound
		}
		c.l.Error().Err(err).Str("module", "books_controller").Str("function", "UpdateBook").Int("book_id", int(bookID)).Msg("error deleting book from service")
		return ctx.Status(responseStatus).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
