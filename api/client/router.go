package routes

import (
	"food-truck-api/api/client/handlers"
	"food-truck-api/package/auth"
	"food-truck-api/package/client"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ClientRouter(
	app fiber.Router,
	service client.Service,
	authService auth.Service,
	validate validator.Validate,
) {

	app.Post("/client/authenticate", handlers.AuthenticateHandler(service, authService, validate))

}
