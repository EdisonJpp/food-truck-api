package routes

import (
	"food-truck-api/api/product/handlers"
	"food-truck-api/package/product"

	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app fiber.Router, service product.Service) {
	app.Get("/products", handlers.GetProductsHandler(service))
}
