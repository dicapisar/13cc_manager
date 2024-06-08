package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type LogoutHandler interface {
}

type LogoutHandlerImpl struct {
	SessionStore *session.Store
}

func (h *LogoutHandlerImpl) logoutGet(c *fiber.Ctx) error {

	userSession, err := h.SessionStore.Get(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userSession.Delete("logged_in")
	userSession.Delete("user_id")

	if err := userSession.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect("/login")

}
