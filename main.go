package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jayanthkrishna/ecommerce-golang/routes"
)

func main() {

	app := fiber.New()

	routes.GetRoutes(app)
	app.Listen(":8085")

}
