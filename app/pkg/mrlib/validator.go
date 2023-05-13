package mrlib

import (
    "auth-le-back/pkg/mrapp"
    "fmt"
    "reflect"
    "strings"

    "github.com/go-playground/validator/v10"
)

type Validator struct {
    logger mrapp.Logger
    validate *validator.Validate
}

// Make sure the Validator conforms with the mrapp.Validator interface
var _ mrapp.Validator = (*Validator)(nil)

func NewValidator(logger mrapp.Logger) *Validator {
    validate := validator.New()

    validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
        if name == "-" {
            return ""
        }

        return name
    })

    return &Validator {
        logger: logger,
        validate: validate,
    }
}

func (v *Validator) Validate(structure any, errors *mrapp.ErrorList) bool {
    err := v.validate.Struct(structure)

    if err == nil {
        return true
    }

    if _, ok := err.(*validator.InvalidValidationError); ok {
        v.logger.Error(err)
        return false
    }

    isValid := true

    for _, errField := range err.(validator.ValidationErrors) {
        *errors = append(*errors, mrapp.ErrorAttribute{
            Id: errField.Field(),
            Value: fmt.Sprintf("Cause of error: %s", errField.Tag()),
        })

        isValid = false

        v.logger.Debug(
            "Namespace: %s\n" +
            "Field: %s\n" +
            "StructNamespace: %s\n" +
            "StructField: %s\n" +
            "Tag: %s\n" +
            "ActualTag: %s\n" +
            "Kind: %v\n" +
            "Type: %v\n" +
            "Value: %v\n" +
            "Param: %s",
            errField.Namespace(),
            errField.Field(),
            errField.StructNamespace(),
            errField.StructField(),
            errField.Tag(),
            errField.ActualTag(),
            errField.Kind(),
            errField.Type(),
            errField.Value(),
            errField.Param(),
        )
    }

    return isValid
}
