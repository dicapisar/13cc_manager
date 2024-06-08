package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/session"
	"strings"
	"time"
)

func AddingMiddlewares(app *fiber.App, sessionStore *session.Store) {
	applyRequestId(app)
	applyLogger(app)
	applySession(app, sessionStore)
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

func applySession(app *fiber.App, sessionStore *session.Store) {

	urlListWithOutSession := []string{"/login", "/register", "/logout"}

	app.Use(func(c *fiber.Ctx) error {

		userSession, err := sessionStore.Get(c)

		if err != nil {
			return err
		}

		url := string(c.Request().URI().Path())

		for _, urlWithoutSession := range urlListWithOutSession {
			if url == urlWithoutSession {
				userSession.Delete("logged_in")
				userSession.Delete("user_id")
				return c.Next()
			}
		}

		if strings.Contains(url, "/public/") {
			return c.Next()
		}

		isLoggedIn := userSession.Get("logged_in")

		if isLoggedIn == nil {
			return c.Redirect("/login", 302)
		}

		return c.Next()
	})
}
