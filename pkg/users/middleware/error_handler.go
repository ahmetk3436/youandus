package middleware

import (
	"youandus/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx) error {
	storage.CreateTopicExchangeAndDeclareQueue()
	storage.SendLogMessage("hata")
	return c.Status(500).JSON(fiber.Map{
		"message": "Something went wrong",
	})
}
