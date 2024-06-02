package main

import (
	"fmt"
	"github.com/dicapisar/13cc_manager/config"
	"github.com/dicapisar/13cc_manager/database"
	"github.com/dicapisar/13cc_manager/handlers"
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

	app.Get("/ping", handlers.PingHandler).Name("ping")

	port := fmt.Sprintf(":%d", envConfig.Server.Port)

	err = app.Listen(port)

	if err != nil {
		panic(err)
	}

}
