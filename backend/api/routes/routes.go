package routes

import (
	"daily-cute-dogs-email/backend/api/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/", controller.CreateSubscribe)

}
