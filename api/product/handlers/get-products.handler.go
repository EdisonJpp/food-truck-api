package handlers

import (
	"food-truck-api/api/product/presenter"
	"food-truck-api/package/product"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetBooks is handler/controller which lists all Books from the BookShop
func GetProductsHandler(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var allProducts []presenter.Product

		products, err := service.GetProducts()

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProducErrorResponse(err))
		}

		for _, prod := range *products {
			var product = presenter.Product{
				ID:   prod.ID,
				Name: prod.Name,
			}

			allProducts = append(allProducts, product)
		}

		return c.JSON(presenter.ProductsSuccessResponse(&allProducts))
	}
}
