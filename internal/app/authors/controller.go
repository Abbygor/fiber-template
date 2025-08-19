package authors

import (
	"fiber-template/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type AuthorsController interface {
	CreateAuthor(ctx *fiber.Ctx) error
	GetAuthorByID(ctx *fiber.Ctx) error
	GetAuthors(ctx *fiber.Ctx) error
	UpdateAuthor(ctx *fiber.Ctx) error
	DeleteAuthor(ctx *fiber.Ctx) error
}

type ControllerAuthors struct {
	authorsService AuthorsService
	l              zerolog.Logger
}

func NewAuthorsController(authorsService AuthorsService, log zerolog.Logger) AuthorsController {
	return &ControllerAuthors{
		authorsService: authorsService,
		l:              log,
	}
}

func (c *ControllerAuthors) CreateAuthor(ctx *fiber.Ctx) error {
	var author models.Author
	if err := ctx.BodyParser(&author); err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "CreateAuthor").
			Msg("error parsin author")
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdAuthor, err := c.authorsService.CreateAuthor(author)
	if err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "CreateAuthor").
			Msg("error creating author from service")
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(createdAuthor)
}

func (c *ControllerAuthors) GetAuthorByID(ctx *fiber.Ctx) error {
	authorIDStr := ctx.Params("id")
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "GetAuthorByID").
			Str("book_id", authorIDStr).
			Msg("error parsing authorID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	author, err := c.authorsService.GetAuthorByID(ctx.Context(), int(authorID))
	if err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "GetAuthorByID").
			Int("book_id", int(authorID)).
			Msg("error getting authors from service")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(author)
}

func (c *ControllerAuthors) GetAuthors(ctx *fiber.Ctx) error {
	authors, err := c.authorsService.GetAuthors()
	if err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "GetAuthors").
			Msg("error getting all authors from service")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(authors)
}

func (c *ControllerAuthors) UpdateAuthor(ctx *fiber.Ctx) error {
	authorIDStr := ctx.Params("id")
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "UpdateAuthor").
			Str("book_id", authorIDStr).
			Msg("error parsing authorID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var author models.Author
	if err := ctx.BodyParser(&author); err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "UpdateAuthor").
			Msg("error parsing author")
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	updatedAuthor, err := c.authorsService.UpdateAuthor(int(authorID), &author)
	if err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "UpdateAuthor").
			Str("book_id", authorIDStr).
			Msg("error updating author from service")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(updatedAuthor)
}

func (c *ControllerAuthors) DeleteAuthor(ctx *fiber.Ctx) error {
	authorIDStr := ctx.Params("id")
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "DeleteAuthor").
			Str("book_id", authorIDStr).
			Msg("error parsing authorID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = c.authorsService.DeleteAuthor(int(authorID))
	responseStatus := fiber.StatusInternalServerError
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			responseStatus = fiber.StatusNotFound
		}
		c.l.Error().
			Err(err).
			Str("module", "authors_controller").
			Str("function", "DeleteAuthor").
			Int("book_id", int(authorID)).
			Msg("error deleting author from service")
		return ctx.Status(responseStatus).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
