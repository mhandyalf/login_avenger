package validators

import (
	"login_avenger/auth"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(user auth.User) error {
	return validate.Struct(user)
}
