package api

import (
	"daily-cute-dogs-email/backend/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	app := fiber.New()
	app.Use(cors.New())
	routes.Router(app)
	app.Listen(":3001")
}
