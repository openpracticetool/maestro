package validators

import (
	"github.com/go-playground/validator"
)

//ValidateStruct ::: validate the workspace model struct
func ValidateStruct(model interface{}) error {
	validate := validator.New()
	return validate.Struct(model)
}
