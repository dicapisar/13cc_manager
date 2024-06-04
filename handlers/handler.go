package handlers

import "github.com/gofiber/fiber/v2"

func AddingHandlers(app *fiber.App) error {

	app.Get("/ping", PingHandler).Name("ping handler")

	return nil
}
