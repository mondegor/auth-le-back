package dto

import (
    "regexp"

    "github.com/go-playground/validator/v10"
)

var (
    regexpLogin = regexp.MustCompile("^[a-z][a-z0-9]+$")
)

func ValidateLogin() any {
    return func (fl validator.FieldLevel) bool {
        return regexpLogin.MatchString(fl.Field().String())
    }
}
