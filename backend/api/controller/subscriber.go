package controller

import (
	"daily-cute-dogs-email/backend/api/service"
	"daily-cute-dogs-email/backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateSubscribe(c *fiber.Ctx) error {
	var subscriber models.Subscriber
	c.BodyParser(&subscriber)
	if err := service.CreateSubscribe(subscriber.Email); err != nil {
		c.Status(400).JSON(fiber.Map{"error": err})
		return err
	} else {
		c.Status(201).JSON(subscriber)
		return nil
	}
}
