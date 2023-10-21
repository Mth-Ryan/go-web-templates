package controllers

import (
	"time"

	"github.com/Mth-Ryan/waveaction/pkg/application/dtos"
	"github.com/Mth-Ryan/waveaction/pkg/application/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BooksController struct{
	validator interfaces.JsonValidator
}

func NewBooksController(
	validator interfaces.JsonValidator,
) *BooksController {
	return &BooksController{
		validator,
	}
}

func (bc *BooksController) GetAll(ctx *fiber.Ctx) error {
	return ctx.JSON([]dtos.BookOutputDto{
		{
			ID:        uuid.New(),
			Title:     "Game of Thrones",
			Author:    "J.R.R Martin",
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Title:     "Fire and Blood",
			Author:    "J.R.R Martin",
			CreatedAt: time.Now(),
		},
	})
}

func (bc *BooksController) Get(ctx *fiber.Ctx) error {
	return ctx.JSON(dtos.BookOutputDto{
		ID:        uuid.New(),
		Title:     "Game of Thrones",
		Author:    "J.R.R Martin",
		CreatedAt: time.Now(),
	})
}

func (bc *BooksController) Create(ctx *fiber.Ctx) error {
	input := new(dtos.BookInputDto)
	ctx.BodyParser(input)

	ok, errors := bc.validator.Validate(input)
	if (!ok) {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return ctx.JSON(dtos.BookOutputDto{
		ID:        uuid.New(),
		Title:     input.Title,
		Author:    input.Author,
		CreatedAt: time.Now(),
	})
}

func (bc *BooksController) Update(ctx *fiber.Ctx) error {
	input := new(dtos.BookInputDto)
	ctx.BodyParser(input)

	ok, errors := bc.validator.Validate(input)
	if (!ok) {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return ctx.JSON(dtos.BookOutputDto{
		ID:        uuid.New(),
		Title:     input.Title,
		Author:    input.Author,
		CreatedAt: time.Now(),
	})
}

func (bc *BooksController) Delete(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func (bc *BooksController) GetRouter(app *fiber.App) fiber.Router {
	router := app.Group("/books")

	router.Get("/", bc.GetAll)
	router.Get("/:id", bc.Get)
	router.Post("/", bc.Create)
	router.Put("/:id", bc.Update)
	router.Delete("/:id", bc.Delete)

	return router
}
