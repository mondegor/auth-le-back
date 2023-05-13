package mrlib

import (
    "auth-le-back/pkg/mrapp"
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

func (v *Validator) Register(tag string, fn mrapp.ValidatorTagFunc) {
    var err error

    if vfn, ok := fn().(func (fl validator.FieldLevel) bool); ok {
        err = v.validate.RegisterValidation(tag, vfn)

        if err != nil {
            v.logger.Error(err)
        }
    } else {
        v.logger.Error(mrapp.ErrInternalTypeAssertion)
    }
}

func (v *Validator) Validate(structure any) error {
    err := v.validate.Struct(structure)

    if err == nil {
        return nil
    }

    if _, ok := err.(*validator.InvalidValidationError); ok {
        return mrapp.ErrInternal.Wrap(err)
    }

    errors := &mrapp.ErrorList{}

    for _, errField := range err.(validator.ValidationErrors) {
        errors.Add(errField.Field(), "Cause of error: %s", errField.Tag())

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

    return errors
}
