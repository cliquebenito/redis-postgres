package controller

import (
	"CRUD-REDIS/domain"
	"CRUD-REDIS/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type NovelController struct {
	novelUsecase domain.NovelUsecase
}

func NewNovelController(novelUsecase domain.NovelUsecase) *NovelController {
	return &NovelController{novelUsecase: novelUsecase}
}

func (n *NovelController) CreateNovel(ctx *fiber.Ctx) error {
	var novelRequest model.Novel
	var resposne model.Response
	if err := ctx.BodyParser(&novelRequest); err != nil {
		resposne = model.Response{
			StatusCode: 400,
			Message:    "Bad request",
		}
		return ctx.Status(400).JSON(resposne)
	}

	if novelRequest.Author == "" && novelRequest.Name == "" && novelRequest.Description == "" {
		resposne = model.Response{
			StatusCode: 400,
			Message:    "Bad request",
		}
		return ctx.Status(400).JSON(resposne)

	}
	if err := n.novelUsecase.CreateNovel(novelRequest); err != nil {
		resposne = model.Response{
			StatusCode: 400,
			Message:    "Bad Request",
		}
		return ctx.Status(400).JSON(resposne)
	}
	resposne = model.Response{
		StatusCode: http.StatusInternalServerError,
		Message:    "Insert Success",
	}

	return ctx.Status(http.StatusInternalServerError).JSON(resposne)
}
