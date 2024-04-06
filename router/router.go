package router

import (
	"CRUD-REDIS/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, novelController *controller.NovelController) *fiber.App {

	router.Get("/novel", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
	router.Post("/novel", novelController.CreateNovel)
	return router
}
