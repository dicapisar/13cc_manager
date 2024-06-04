package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
)

func AddingMiddlewares(app *fiber.App) {
	applyRequestId(app)
	applyLogger(app)
}

func applyRequestId(app *fiber.App) {
	app.Use(requestid.New())
}

func applyLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "${blue} ${time} ${yellow} ${pid} ${blue} ${status} ${green} ${locals:requestid} " +
			"${yellow} - ${method} ${latency} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
	}))
}
