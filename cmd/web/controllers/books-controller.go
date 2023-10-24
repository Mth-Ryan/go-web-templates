package controllers

import (
	"fmt"

	"github.com/Mth-Ryan/waveaction/cmd/web/views"
	"github.com/Mth-Ryan/waveaction/internal/application/dtos"
	"github.com/Mth-Ryan/waveaction/internal/application/interfaces"
	"github.com/Mth-Ryan/waveaction/internal/application/services"
	"github.com/gofiber/fiber/v2"
)

type BooksController struct{
	validator interfaces.JsonValidator
	service services.BooksService
	views views.ViewsRenderer
}

func NewBooksController(
	validator interfaces.JsonValidator,
	service services.BooksService,
	views views.ViewsRenderer,
) *BooksController {
	return &BooksController{
		validator,
		service,
		views,
	}
}

func (bc *BooksController) GetAll(ctx *fiber.Ctx) error {
	books, err := bc.service.GetAll()
	if err != nil {
		return err
	}

	return renderView(
		ctx,
		bc.views,
		"./templates/books/index.tmpl.html",
		map[string]any{
			"title": "Books",
			"books": books,
		},
	)
}

func (bc *BooksController) Get(ctx *fiber.Ctx) error {
	id, err := bindUUIDParam(ctx, "id")
	if (err != nil) {
		return err
	}

	book, err := bc.service.Get(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return renderView(
		ctx,
		bc.views,
		"./templates/books/show.tmpl.html",
		map[string]any{
			"title": "Books",
			"book": book,
		},
	)
}

func (bc *BooksController) Create(ctx *fiber.Ctx) error {
	return renderView(
		ctx,
		bc.views,
		"./templates/books/books-form.tmpl.html",
		map[string]any{
			"title": "Books",
			"variantTitle": "Create",
		},
	)
}

func (bc *BooksController) CreateSubmit(ctx *fiber.Ctx) error {
	input := new(dtos.BookInputDto)
	if err := bindAndValidate(ctx, bc.validator, input); err != nil {
		return err
	}

	book, err := bc.service.Create(*input)
	if (err != nil) {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	
	return ctx.Redirect(fmt.Sprintf("/books/%s", book.ID.String()), 302)
}

func (bc *BooksController) Update(ctx *fiber.Ctx) error {
	id, err := bindUUIDParam(ctx, "id")
	if (err != nil) {
		return err
	}

	book, err := bc.service.Get(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return renderView(
		ctx,
		bc.views,
		"./templates/books/books-form.tmpl.html",
		map[string]any{
			"title": "Books",
			"variantTitle": "Edit",
			"book": book,
		},
	)
}

func (bc *BooksController) UpdateSubmit(ctx *fiber.Ctx) error {
	id, err := bindUUIDParam(ctx, "id")
	if (err != nil) {
		return err
	}

	input := new(dtos.BookInputDto)
	if err := bindAndValidate(ctx, bc.validator, input); err != nil {
		return err
	}

	book, err := bc.service.Update(id, *input)
	if (err != nil) {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Redirect(fmt.Sprintf("/books/%s", book.ID.String()), 302)
}

func (bc *BooksController) Delete(ctx *fiber.Ctx) error {
	id, err := bindUUIDParam(ctx, "id")
	if (err != nil) {
		return err
	}

	err = bc.service.Delete(id)
	if (err != nil) {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (bc *BooksController) RegisterController(app *fiber.App) {
	router := app.Group("/books")

	router.Get("/", bc.GetAll)
	router.Get("/create", bc.Create)
	router.Post("/create", bc.CreateSubmit)
	router.Get("/:id", bc.Get)
	router.Get("/:id/edit", bc.Update)
	router.Post("/:id/edit", bc.UpdateSubmit)
}
