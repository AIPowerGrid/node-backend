package validators

import (
	"backend/api"
	"backend/core"

	"github.com/go-playground/validator/v10"
)

var (
	log      = core.GetLogger()
	validate = validator.New()
)

func parseErrs(err error) []*api.ErrorResponse {
	var errors []*api.ErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		var element api.ErrorResponse
		// element.FailedField = err.StructNamespace()
		element.FailedField = err.Field()
		// element.Tag = err.Tag()
		element.Value = err.Param()
		errors = append(errors, &element)
	}
	return errors
}

func ValidateInterface(data interface{}) []*api.ErrorResponse {
	var errors []*api.ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		errors = parseErrs(err)
	}
	return errors
}
