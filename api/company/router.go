package company

import (
	"food-truck-api/api/company/handlers"
	"food-truck-api/package/company/contract"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CompanyRouter(app fiber.Router, service contract.Service, validate validator.Validate) {
	app.Post("/company/register", handlers.RegisterHandler(service, validate))
}
