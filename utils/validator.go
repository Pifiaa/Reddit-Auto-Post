package utils

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value, omitempty"`
}

var validate = validator.New()

func ValidateStruct[T any](model T) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(model)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)
		}
	}
	return errors
}
