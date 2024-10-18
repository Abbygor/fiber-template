package authors

import (
	"fiber-template/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
}

func NewAuthorsController(authorsService AuthorsService) AuthorsController {
	return &ControllerAuthors{
		authorsService: authorsService,
	}
}

func (c *ControllerAuthors) CreateAuthor(ctx *fiber.Ctx) error {
	var author models.Author
	if err := ctx.BodyParser(&author); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	createdAuthor, err := c.authorsService.CreateAuthor(author)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(createdAuthor)
}

func (c *ControllerAuthors) GetAuthorByID(ctx *fiber.Ctx) error {
	authorIDStr := ctx.Params("id")
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	author, err := c.authorsService.GetAuthorByID(int(authorID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(author)
}

func (c *ControllerAuthors) GetAuthors(ctx *fiber.Ctx) error {
	authors, err := c.authorsService.GetAuthors()
	if err != nil {
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var author models.Author
	if err := ctx.BodyParser(&author); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	updatedAuthor, err := c.authorsService.UpdateAuthor(int(authorID), &author)
	if err != nil {
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
		return ctx.Status(responseStatus).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
