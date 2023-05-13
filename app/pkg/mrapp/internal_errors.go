package mrapp

var (
    ErrSystem = NewError(
        errorCodeSystem, ErrorKindSystem, "system server error")

    ErrInternal = NewError(
        errorCodeInternal, ErrorKindInternal, "internal server error")

    ErrInternalTypeAssertion = NewError(
        "errInternalTypeAssertion", ErrorKindInternal, "type assertion error")
)
