package handlers

import (
	"github.com/dicapisar/13cc_manager/repository"
	"github.com/dicapisar/13cc_manager/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddingHandlers(app *fiber.App, DB *gorm.DB) error {

	app.Get("/ping", pingHandler).Name("ping handler")

	setUsersGroupHandler(app, DB)
	return nil
}

func setUsersGroupHandler(app *fiber.App, DB *gorm.DB) {

	userRepository := repository.UserRepositoryImpl{DB: DB}

	userService := services.UserServiceImpl{UserRepository: &userRepository}

	userHandler := UserHandlerImpl{
		userService: &userService,
	}

	usersRouteGroupHandler := app.Group("/users")
	usersRouteGroupHandler.Post("/", userHandler.createNewUserHandler).Name("create new user handler")
}
