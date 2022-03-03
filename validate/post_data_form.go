package validate

import (
	"gopkg.in/go-playground/validator.v9"
)

type UserPostFormUpload struct {
	Email  string `json:"email" validate:"required,email,passStrong"`
	Name   string `json:"name" validate:"required"`
	Detail string `json:"detail" validate:"required"`
}

func (user *UserPostFormUpload) Validate() error {
	v := validator.New()
	// customs validate
	err := ruleCustoms(v, "passStrong")
	if err != nil {
		return err
	}

	return v.Struct(user)
}
