package validate

import (
	"github.com/go-playground/validator/v10"
)

//v.RegisterValidation("passwd", )

func lengthStrong(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 6
}
func ruleCustoms(v *validator.Validate, ruleNameCustoms string) error {
	switch ruleNameCustoms {
	case "lengthStrong":
		return v.RegisterValidation("lengthStrong", lengthStrong)
	default:
		panic("ruleNameCustoms not exist")
	}
}
