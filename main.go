package main

import (
	"fmt"
	auth2 "github.com/dicapisar/13cc_manager/auth"
	"github.com/dicapisar/13cc_manager/config"
	"github.com/dicapisar/13cc_manager/database"
	"github.com/dicapisar/13cc_manager/handlers"
	"github.com/dicapisar/13cc_manager/middlewares"
	"github.com/dicapisar/13cc_manager/repository"
	"github.com/dicapisar/13cc_manager/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/handlebars/v2"
	"time"
)

func main() {

	envConfig := config.GetConfig()

	dependencies, err := startDependencies(envConfig)

	if err != nil {
		panic(err)
	}

	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		AppName:           envConfig.App.Name,
		CaseSensitive:     envConfig.Server.CaseSensitive,
		EnablePrintRoutes: envConfig.Server.EnablePrintRoutes,
		Views:             engine,
	})

	// Adding Middleware
	middlewares.AddingMiddlewares(app)

	// Adding Static Directory
	app.Static("/public", "./public").Name("static_files")

	// Adding handlers
	err = handlers.AddingHandlers(app, dependencies)
	if err != nil {
		panic(err)
	}

	port := fmt.Sprintf(":%d", envConfig.Server.Port)
	err = app.Listen(port)

	if err != nil {
		panic(err)
	}

}

func startDependencies(config *config.Config) (*handlers.Dependencies, error) {
	db, err := database.NewDatabase(&config.Database)

	if err != nil {
		return nil, err
	}

	userRepository := repository.UserRepositoryImpl{
		DB: db.DB,
	}

	userService := services.UserServiceImpl{UserRepository: &userRepository}
	loginService := services.LoginServiceImpl{UserRepository: &userRepository}

	auth := auth2.Auth{Config: &config.Auth}

	store := session.New(session.Config{
		Expiration:   time.Duration(config.Auth.SessionExpiration),
		CookieSecure: true,
	})

	return &handlers.Dependencies{
		UserService:  &userService,
		LoginService: &loginService,
		Auth:         &auth,
		SessionStore: store,
	}, nil
}
