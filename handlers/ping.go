package handlers

import "github.com/gofiber/fiber/v2"

func PingHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("pong")
}