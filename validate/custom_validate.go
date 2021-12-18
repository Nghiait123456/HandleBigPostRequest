package validate

import (
	"errors"
	"gopkg.in/go-playground/validator.v9"
)

//v.RegisterValidation("passwd", )

func passStrong(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 6
}
func ruleCustoms(v *validator.Validate, ruleNameCustoms string) error {
	switch ruleNameCustoms {
	case "passStrong":
		return v.RegisterValidation("passStrong", passStrong)
	default:
		return errors.New("ruleNameCustoms not exist")
	}
}


