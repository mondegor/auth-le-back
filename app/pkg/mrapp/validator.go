package mrapp

import "fmt"

type ErrorAttribute struct {
    Id string `json:"id"`
    Value string `json:"value"`
}

type ErrorList []ErrorAttribute

func (e *ErrorList) Add(id, value string, args ...any) {
    if len(args) > 0 {
        value = fmt.Sprintf(value, args...)
    }

    *e = append(*e, ErrorAttribute{Id: id, Value: value})
}

type Validator interface {
    Validate(structure any, errors *ErrorList) bool
}
