package mrapp

import "errors"

var (
    ErrStorageConnectionAlreadyExists = NewError(
        "errStorageConnectionAlreadyExists", errors.New("connection already exists"), ErrorKindSystem)

    ErrStorageConnectionFailed = NewError(
        "errStorageConnectionFailed", errors.New("connection is failed"), ErrorKindSystem)

    ErrStorageQueryFailed = NewError(
        "errStorageQueryFailed", errors.New("query is failed"), ErrorKindInternal)

    ErrStorageFetchDataFailed = NewError(
        "errStorageFetchDataFailed", errors.New("fetching data is failed"), ErrorKindInternal)

    ErrStorageNoRowFound = NewError(
        "errStorageNoRowFound", errors.New("no row found"), ErrorKindSystem)
)
