package main

import (
	"fmt"
	"github.com/dicapisar/13cc_manager/config"
	"github.com/dicapisar/13cc_manager/database"
	"github.com/dicapisar/13cc_manager/handlers"
	"github.com/dicapisar/13cc_manager/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {

	envConfig := config.GetConfig()

	newDatabase, err := database.NewDatabase(&envConfig.Database)

	if err != nil {
		panic(err)
	}

	fmt.Println(newDatabase)

	app := fiber.New(fiber.Config{
		AppName:           envConfig.App.Name,
		CaseSensitive:     envConfig.Server.CaseSensitive,
		EnablePrintRoutes: envConfig.Server.EnablePrintRoutes,
	})

	// Adding Middleware
	middlewares.AddingMiddlewares(app)

	// Adding Static Directory
	app.Static("/public", "./public").Name("static files")

	// Adding handlers
	err = handlers.AddingHandlers(app, newDatabase.DB)
	if err != nil {
		panic(err)
	}

	port := fmt.Sprintf(":%d", envConfig.Server.Port)
	err = app.Listen(port)

	if err != nil {
		panic(err)
	}

}
