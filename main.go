package main

import (
	"fmt"
	"github.com/dicapisar/13cc_manager/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

	envConfig := config.GetConfig()

	app := fiber.New(fiber.Config{
		AppName:           envConfig.App.Name,
		CaseSensitive:     envConfig.Server.CaseSensitive,
		EnablePrintRoutes: envConfig.Server.EnablePrintRoutes,
	})

	port := fmt.Sprintf(":%d", envConfig.Server.Port)

	err := app.Listen(port)

	if err != nil {
		panic(err)
	}

}
