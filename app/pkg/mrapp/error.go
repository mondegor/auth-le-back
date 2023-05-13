package mrapp

import "fmt"

const (
    ErrorKindInternal = 1
    ErrorKindSystem   = 2
    ErrorKindUser     = 3
)

type ErrorKind int32

type AppError struct {
    code string
    err  error
    kind ErrorKind
}

func NewError(code string, err error, kind ErrorKind) *AppError {
    return &AppError{
        code: code,
        err:  err,
        kind: kind,
    }
}

// Wrap it clones the current structure and binds the passed error
func (e *AppError) Wrap(err error) error {
    if err == nil {
        panic("error is nil, wrapping is not necessary")
    }

    k := *e
    k.err = err

    return &k
}

func (e *AppError) Code() string {
    return e.code
}

func (e *AppError) Kind() ErrorKind {
    return e.kind
}

func (e *AppError) Error() string {
    return fmt.Sprintf("%s: %v", e.code, e.err)
}

func (e *AppError) Unwrap() error {
    return e.err
}
