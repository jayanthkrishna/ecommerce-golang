package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jayanthkrishna/ecommerce-golang/controllers"
)

func GetRoutes(app *fiber.App) {

	app.Post("/users/signup", controllers.SignUp())
	app.Post("/users/login", controllers.Login())
	app.Post("/admin/addproduct", controllers.ProductViewerAdmin())
	app.Get("/users/productview", controllers.SearchProduct())
	app.Get("/users/search", controllers.SearchProductByQuery())

}
