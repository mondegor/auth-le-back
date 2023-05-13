package mrapp

var (
    ErrHttpRequestParseData = NewError(
        "errHttpRequestParseData", ErrorKindUser, "request body is not valid")

    ErrHttpResponseParseData = NewError(
        "errHttpResponseParseData", ErrorKindInternal, "response data is not valid")

    ErrHttpResponseSendData = NewError(
        "errHttpResponseSendData", ErrorKindInternal, "response data is not send")

    ErrHttpResponseSystemTemporarilyUnableToProcess = NewError(
       "errHttpResponseSystemTemporarilyUnableToProcess", ErrorKindUser, "the system is temporarily unable to process your request. Please try again")

    ErrHttpResourceNotFound = NewError(
        "errHttpResuorceNotFound", ErrorKindUser, "resource not found")
)
