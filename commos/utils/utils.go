package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetUserDataFromSessionStorage(sessionStore *session.Store, c *fiber.Ctx) (fiber.Map, error) {
	userSession, err := sessionStore.Get(c)

	if err != nil {
		return nil, err
	}

	userId := userSession.Get("user_id")
	userName := userSession.Get("user_name")
	userRol := userSession.Get("user_rol")
	userLoggedIn := userSession.Get("logged_in")

	return fiber.Map{
		"userId":   userId,
		"userName": userName,
		"userRol":  userRol,
		"loggedIn": userLoggedIn,
	}, nil
}
