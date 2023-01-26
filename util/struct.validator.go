package util

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/medical-identification/go-std/model"
)

var validate = validator.New()

func init() {
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{}, sql.NullInt16{}, sql.NullTime{}, sql.NullByte{}, sql.NullInt32{}, uuid.NullUUID{})
}

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

// ValidateValuer implements validator.CustomTypeFunc
func ValidateValuer(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
}
