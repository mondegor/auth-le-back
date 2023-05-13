package mrapp

import "errors"

var (
    ErrServiceResourceTemporarilyUnavailable = NewError(
        "errServiceResourceTemporarilyUnavailable", errors.New("resource is temporarily unavailable"), ErrorKindUser)

    ErrServiceResourceNotFound = NewError(
        "errServiceResourceNotFound", errors.New("resource is not found"), ErrorKindUser)

    ErrServiceResourceNotCreated = NewError(
        "errServiceResourceNotCreated", errors.New("resource is not created"), ErrorKindUser)

    ErrServiceResourceNotUpdated = NewError(
        "errServiceResourceNotUpdated", errors.New("resource is not updated"), ErrorKindUser)

    ErrServiceResourceNotRemoved = NewError(
        "errServiceResourceNotRemoved", errors.New("resource is not removed"), ErrorKindUser)

    ErrServiceIncorrectData = NewError(
        "errServiceIncorrectData", errors.New("data is incorrect"), ErrorKindUser)
)
