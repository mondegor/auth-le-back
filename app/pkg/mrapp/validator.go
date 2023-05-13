package mrapp

import "fmt"

type (
    ErrorAttribute struct {
        Id string `json:"id"`
        Value string `json:"value"`
    }

    ErrorList []ErrorAttribute

    ValidatorTagFunc func() any
)

func (e *ErrorList) Error() string {
    return fmt.Sprintf("%+v", *e)
}

func (e *ErrorList) Add(id, value string, args ...any) {
    if len(args) > 0 {
        value = fmt.Sprintf(value, args...)
    }

    *e = append(*e, ErrorAttribute{Id: id, Value: value})
}

type Validator interface {
    Register(tag string, fn ValidatorTagFunc)
    Validate(structure any) error
}
