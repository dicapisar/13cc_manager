package handlers

import "github.com/gofiber/fiber/v2"

func pingHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("pong")
}
