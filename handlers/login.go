package handlers

import (
	"github.com/dicapisar/13cc_manager/dtos/request"
	"github.com/dicapisar/13cc_manager/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type LoginHandler interface {
}

type LoginHandlerImpl struct {
	LoginService   *services.LoginServiceImpl
	UserRolService *services.UserRolServiceImpl
	SessionStore   *session.Store
}

func (h *LoginHandlerImpl) loginGet(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func (h *LoginHandlerImpl) loginPost(c *fiber.Ctx) error {

	var req request.Login

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.LoginService.Login(req.Email, req.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userSession, err := h.SessionStore.Get(c)

	if err != nil {
		// TODO: CREAR VISTA PARA ERRORES 404 Y DE SERVIDOR
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userRol, err := h.UserRolService.GetUserRolByID(int(user.ID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	userSession.Set("logged_in", true)
	userSession.Set("user_id", user.ID)
	userSession.Set("user_name", user.Name)
	userSession.Set("user_rol", userRol.Rol.Name)

	if err := userSession.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect("/home")
}
