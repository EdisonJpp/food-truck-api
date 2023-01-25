package handlers

import (
	"food-truck-api/api/company/presenter"
	"food-truck-api/package/company/contract"
	"net/http"

	sPresenters "food-truck-api/shared/presenter"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(service contract.Service, validate validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(contract.RegisterRequest)

		err := c.BodyParser(requestBody)

		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
		}

		if err := validate.Struct(requestBody); err != nil {
			errors := sPresenters.FieldErrorFormatPresenter(err.(validator.ValidationErrors))
			return c.Status(fiber.StatusBadRequest).JSON(sPresenters.ErrorResponsePresenter(errors))
		}

		registed, err := service.Register(requestBody)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
		}

		c.Status(http.StatusCreated)
		return c.JSON(presenter.RegisterPresent(registed, "one-access-token"))
	}
}
