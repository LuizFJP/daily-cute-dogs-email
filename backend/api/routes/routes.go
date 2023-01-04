package routes

import (
	"daily-cute-dogs-email/backend/api/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/", controller.CreateSubscribe)
	api.Delete("/", controller.DeleteSubscribe)

}
