package routes

import (
	"food-truck-api/api/product/handlers"
	"food-truck-api/package/product"
	"food-truck-api/shared/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app fiber.Router, service product.Service) {
	app.Get("/products", middleware.Protected(), handlers.GetProductsHandler(service))
}
