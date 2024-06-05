package handlers

import (
	auth2 "github.com/dicapisar/13cc_manager/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type HomeHandler interface {
}

type HomeHandlerImpl struct {
	SessionStore *session.Store
	Auth         *auth2.Auth
}

func (h *HomeHandlerImpl) home(c *fiber.Ctx) error {

	userSession, err := h.SessionStore.Get(c)

	if !h.Auth.IsLoggedIn(userSession) || err != nil {
		return c.Redirect("/login", 302)
	}

	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	}, "layouts/main")
}

func (h *HomeHandlerImpl) main(c *fiber.Ctx) error {
	return c.Redirect("/home", 302)
}
