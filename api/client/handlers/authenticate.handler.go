package handlers

import (
	"food-truck-api/api/client/contract"
	"food-truck-api/api/client/presenter"
	"food-truck-api/package/auth"
	authContract "food-truck-api/package/auth/contract"
	"food-truck-api/package/client"
	"food-truck-api/package/entities"
	sPresenters "food-truck-api/shared/presenter"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthenticateHandler(
	service client.Service,
	authService auth.Service,
	validate validator.Validate,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(contract.AuthenticateRequest)
		err := c.BodyParser(requestBody)

		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
		}

		if err := validate.Struct(requestBody); err != nil {
			errors := sPresenters.FieldErrorFormatPresenter(err.(validator.ValidationErrors))
			return c.Status(fiber.StatusBadRequest).JSON(sPresenters.ErrorResponsePresenter(errors))
		}

		client := new(entities.Client)

		if requestBody.Email != "" {

			isEmailExists, err := service.IsEmailExists(requestBody.Email, requestBody.CompanyId)

			if err != nil {
				c.Status(http.StatusInternalServerError)
				return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
			}

			if isEmailExists {
				item, err := service.GetClientByEmail(requestBody.Email, requestBody.CompanyId)

				if err != nil {

					if err == mongo.ErrNoDocuments {
						c.Status(http.StatusNotFound)
						return c.JSON(sPresenters.ErrorResponsePresenter("Client Not Found"))
					}

					c.Status(http.StatusInternalServerError)
					return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
				}

				client.ID = item.ID
				client.Name = item.Name
				client.Email = item.Email

			} else {

				client.Name = requestBody.Name
				client.Email = requestBody.Email

				newClient, err := service.CreateClient(client)

				if err != nil {
					c.Status(http.StatusInternalServerError)
					return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
				}

				client.ID = newClient.ID
			}
		}

		t, err := authService.CreateToken(&authContract.CreateTokenRequest{
			ID:    client.ID,
			Name:  client.Name,
			Email: client.Email,
		})

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(sPresenters.ErrorResponsePresenter(err.Error()))
		}

		c.Status(http.StatusCreated)
		return c.JSON(presenter.AuthenticatePresent(&t, client))
	}
}
