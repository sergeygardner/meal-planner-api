package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var service *validator.Validate

func init() {
	service = validator.New()
}

func Validate(entity any) error {
	errorStruct := service.Struct(entity)

	if errorStruct != nil {
		return errors.Wrapf(errorStruct, "an error occured while validation %v", entity)
	}

	return nil
}
