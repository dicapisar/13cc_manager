package handlers

import (
	auth2 "github.com/dicapisar/13cc_manager/auth"
	"github.com/dicapisar/13cc_manager/commos/utils"
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

	userData, err := utils.GetUserDataFromSessionStorage(h.SessionStore, c)

	if err != nil {
		// TODO: CREAR HANDLER DE ERRORES 500 Y 400
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userData["title"] = "Dashboard"

	return c.Render("index", userData, "layouts/main")
}

func (h *HomeHandlerImpl) main(c *fiber.Ctx) error {
	return c.Redirect("/home", 302)
}
