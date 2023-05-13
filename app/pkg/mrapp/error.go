package mrapp

import (
    "fmt"
    "runtime"
)

const (
    ErrorKindInternal = 1
    ErrorKindSystem   = 2
    ErrorKindUser     = 3

    errorCodeSystem = "errSystem"
    errorCodeInternal = "errInternal"
)

type ErrorKind int32
type ErrorCode string

type AppErrorTemplate struct {
    code ErrorCode
    kind ErrorKind
    defaultMessage string
}

type AppError struct {
    base *AppErrorTemplate
    err  error
    args []any
    file string
    line int
}

func NewError(code ErrorCode, kind ErrorKind, defaultMessage string) *AppErrorTemplate {
    return &AppErrorTemplate{
        code: code,
        kind: kind,
        defaultMessage: defaultMessage,
    }
}

func (e *AppErrorTemplate) New(args ...any) error {
    this := newError(e, nil, e.kind != ErrorKindUser)

    if len(args) > 0 {
        this.args = args
    }

    return this
}

// Wrap it clones the current structure and binds the passed error
func (e *AppErrorTemplate) Wrap(err error) error {
    if err == nil {
        panic("error is nil, wrapping is not necessary")
    }

    return newError(e, err, e.kind != ErrorKindUser)
}

func newError(base *AppErrorTemplate, err error, useCaller bool) *AppError {
    this := &AppError{
        base: base,
        err:  err,
    }

    if useCaller {
        _, file, line, ok := runtime.Caller(2)

        if ok {
            this.file = file
            this.line = line
        }
    }

    return this
}

func (e *AppErrorTemplate) Is(err error) bool {
  if v, ok := err.(*AppErrorTemplate); ok && e.code == v.code {
      return true
  }

  return false
}

func (e *AppErrorTemplate) Error() string {
    return e.defaultMessage
}

//func (e *AppError) Code() ErrorCode {
//    return e.code
//}
//
//func (e *AppError) Kind() ErrorKind {
//    return e.kind
//}

func (e *AppError) New(args ...any) error {
    return e.base.New(args)
}

func (e *AppError) Wrap(err error) error {
    return e.base.Wrap(err)
}

func (e *AppError) Unwrap() error {
    return e.err
}

func (e *AppError) Is(err error) bool {
    if e.base.Is(err) {
        return true
    }

    if v, ok := err.(*AppError); ok && e.base.code == v.base.code {
        return true
    }

    return false
}

func (e *AppError) Error() string {
    errMessage := e.base.defaultMessage

    if len(e.args) > 0 {
        errMessage = fmt.Sprintf(errMessage, e.args)
    }

    if e.err == nil {
        if len(e.file) == 0 {
            return errMessage
        }

        return fmt.Sprintf("%s in %s:%d", errMessage, e.file, e.line)
    }

    if len(e.file) == 0 {
        return fmt.Sprintf("%s: %s", errMessage, e.err.Error())
    }

    return fmt.Sprintf("%s in %s:%d; %s", errMessage, e.file, e.line, e.err.Error())
}

func (e *AppError) UserError(locale Locale) ErrorMessage {
    if e.base.kind != ErrorKindInternal {
        // :TODO: можно добавить именованные аргументы
        mess := locale.GetError(e.base.code)

        if mess.Reason == "" {
            mess.Reason = e.base.defaultMessage
        }

        if len(e.args) > 0 {
            mess.Reason = fmt.Sprintf(mess.Reason, e.args...)
        }

        return mess
    }

    return locale.GetError(errorCodeInternal)
}
