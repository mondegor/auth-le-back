package usecase

import (
    "auth-le-back/pkg/mrapp"
)

var (
    ErrAuthUserLoginAlreadyExists = mrapp.NewError(
        "errAuthUserLoginAlreadyExists", mrapp.ErrorKindUser, "Login '%s' is already exists")

    ErrAuthUserEmailAlreadyExists = mrapp.NewError(
        "errAuthUserEmailAlreadyExists", mrapp.ErrorKindUser, "Email '%s' is already exists")
)
