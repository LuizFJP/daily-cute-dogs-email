package api

import (
	"daily-cute-dogs-email/backend/api/routes"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()
	routes.Router(app)
	app.Listen(":3001")
}
