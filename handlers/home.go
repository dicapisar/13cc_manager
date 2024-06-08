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

	if err != nil {
		// TODO: CREAR HANDLER DE ERRORES 500 Y 400
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userId := userSession.Get("user_id")
	userName := userSession.Get("user_name")
	userRol := userSession.Get("user_rol")

	return c.Render("index", fiber.Map{
		"userId":   userId,
		"userName": userName,
		"userRol":  userRol,
	}, "layouts/main")
}

func (h *HomeHandlerImpl) main(c *fiber.Ctx) error {
	return c.Redirect("/home", 302)
}
