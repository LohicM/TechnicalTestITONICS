package main

import (
	"TechnicalTestITONICS/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//routes
	routes.UserRoute(app) //add this

	app.Listen(":5000")
}
