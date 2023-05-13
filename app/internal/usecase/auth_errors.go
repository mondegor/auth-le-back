package usecase

import (
    "auth-le-back/pkg/mrapp"
    "errors"
)

var (
    AuthErrUserEmailAlreadyExists = mrapp.NewError(
        "authErrUserEmailAlreadyExists", errors.New("email is already exists"), mrapp.ErrorKindUser)
)
