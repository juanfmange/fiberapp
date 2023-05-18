package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	config "github.com/juanfmange/fiberapp/Config"
	middleware "github.com/juanfmange/fiberapp/Middleware"
	routes "github.com/juanfmange/fiberapp/Routes"
)

func main() {
	middleware.LoadEnv()
	config.Connection()

	app := fiber.New()
	// app.Use(app)

	routes.Setup(app)

	app.Listen(os.Getenv("PORT"))

}
