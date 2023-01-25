package handlers

import (
	"food-truck-api/api/company/presenter"
	"food-truck-api/package/auth"
	authContract "food-truck-api/package/auth/contract"
	"food-truck-api/package/company/contract"
	"net/http"

	sPresenters "food-truck-api/shared/presenter"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginHandler(service contract.Service, authService auth.Service, validate validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(contract.LoginRequest)

		err := c.BodyParser(requestBody)

		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
		}

		if err := validate.Struct(requestBody); err != nil {
			errors := sPresenters.FieldErrorFormatPresenter(err.(validator.ValidationErrors))
			return c.Status(fiber.StatusBadRequest).JSON(sPresenters.ErrorResponsePresenter(errors))
		}

		company, err := service.GetCompanyByEmail(requestBody.Email)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.Status(http.StatusUnauthorized)
				return c.JSON(sPresenters.ErrorResponsePresenter("Incorrect data"))
			}

			c.Status(http.StatusInternalServerError)
			return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
		}

		checkedPassword := authService.CheckPassword(company.Password, requestBody.Password)

		if checkedPassword != true {
			c.Status(http.StatusUnauthorized)
			return c.JSON(sPresenters.ErrorResponsePresenter("Incorrect data"))
		}

		token, err := authService.CreateToken(&authContract.CreateTokenRequest{
			ID:    company.ID,
			Name:  company.Name,
			Email: company.Email,
		})

		if err != nil {

			c.Status(http.StatusInternalServerError)
			return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
		}

		c.Status(http.StatusCreated)
		return c.JSON(presenter.LoginPresent(&token))
	}
}
