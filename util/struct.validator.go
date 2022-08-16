package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/medical-identification/model"
)

var validate = validator.New()

func ValidateStruct(obj interface{}) model.ErrorValidationResponse {
	var errors model.ErrorValidationResponse = model.ErrorValidationResponse{}
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorValidatorMeta
			//element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()

			//errors = append(errors, &element)
			errors[strings.ToLower(err.StructField())] = element
		}
	}
	return errors
}
