package handlers

import (
	"github.com/dicapisar/13cc_manager/dtos/request"
	"github.com/dicapisar/13cc_manager/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	CreateNewUserHandler(c *fiber.Ctx) error
}

type UserHandlerImpl struct {
	userService *services.UserServiceImpl
}

func (u *UserHandlerImpl) createNewUserHandler(c *fiber.Ctx) error {
	var req request.CreateNewUser

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := u.userService.CreateNewUser(&req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
