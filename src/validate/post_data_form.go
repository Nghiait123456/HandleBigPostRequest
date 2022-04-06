package validate

import (
	"github.com/go-playground/validator/v10"
)

type PostFormUpload struct {
	Email  string `json:"email" validate:"required,email,lengthStrong"`
	Name   string `json:"name" validate:"required"`
	Detail string `json:"detail" validate:"required"`
}

func (user *PostFormUpload) Validate() error {
	v := validator.New()
	// customs validate
	err := ruleCustoms(v, "lengthStrong")
	if err != nil {
		return err
	}

	return v.Struct(user)
}
