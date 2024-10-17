package books

import (
	"fiber-template/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
}

func NewBooksController(booksService BooksService) BooksController {
	return &ControllerBooks{
		booksService: booksService,
	}
}

func (c *ControllerBooks) CreateBook(ctx *fiber.Ctx) error {
	var book models.Book
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdBook, err := c.booksService.CreateBook(book)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(createdBook)
}

func (c *ControllerBooks) GetBookByID(ctx *fiber.Ctx) error {
	bookIDStr := ctx.Params("id")
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	book, err := c.booksService.GetBookByID(int(bookID))
	if err != nil {
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	books, err := c.booksService.GetBooksByAuthorID(int(authorID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(books)
}
func (c *ControllerBooks) GetBooks(ctx *fiber.Ctx) error {
	books, err := c.booksService.GetBooks()
	if err != nil {
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var book models.Book
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	book.BookID = int(bookID)

	updatedBook, err := c.booksService.UpdateBook(&book)
	if err != nil {
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
		return ctx.Status(responseStatus).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
