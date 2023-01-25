package presenter

import (
	"github.com/go-playground/validator/v10"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

func FieldErrorFormatPresenter(currentErrors validator.ValidationErrors) []*IError {
	var errors []*IError

	for _, err := range currentErrors {
		var el IError
		el.Field = err.Field()
		el.Tag = err.Tag()
		el.Value = err.Param()
		errors = append(errors, &el)
	}

	return errors

}
