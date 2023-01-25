package presenter

import "github.com/gofiber/fiber/v2"

func ErrorResponsePresenter[T []*IError | string](errors T) *fiber.Map {
	return &fiber.Map{
		"errors": errors,
	}
}
