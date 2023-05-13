package mrapp

var (
    ErrStorageConnectionAlreadyExists = NewError(
        "errStorageConnectionAlreadyExists", ErrorKindSystem, "connection already exists")

    ErrStorageConnectionFailed = NewError(
        "errStorageConnectionFailed", ErrorKindSystem, "connection is failed")

    ErrStorageQueryFailed = NewError(
        "errStorageQueryFailed", ErrorKindInternal, "query is failed")

    ErrStorageFetchDataFailed = NewError(
        "errStorageFetchDataFailed", ErrorKindInternal, "fetching data is failed")

    ErrStorageNoRowFound = NewError(
        "errStorageNoRowFound", ErrorKindUser, "no row found")
)
