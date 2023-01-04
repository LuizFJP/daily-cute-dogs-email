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
		c.Status(400).JSON(err.Error())
		return err
	} else {
		c.Status(201).JSON("email adicionado com sucesso!")
		return nil
	}
}

func DeleteSubscribe(c *fiber.Ctx) error {
	var subscriber models.Subscriber
	c.BodyParser(&subscriber)
	if err := service.DeleteSubscribe(subscriber.Email); err != nil {
		c.Status(400).JSON(err.Error())
		return err
	} else {
		c.Status(200).JSON("email deletado com sucesso! Sentiremos sua falta :(")
		return nil
	}
}
