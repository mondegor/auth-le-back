package mrapp

import "errors"

var (
    ErrHttpRequestParseData = NewError(
        "errHttpRequestParseData", errors.New("request body is not valid"), ErrorKindUser)

    ErrHttpResponseParseData = NewError(
        "errHttpResponseParseData", errors.New("response data is not valid"), ErrorKindInternal)

    ErrHttpResponseSystemTemporarilyUnableToProcess = NewError(
       "errHttpResponseSystemTemporarilyUnableToProcess", errors.New("the system is temporarily unable to process your request. Please try again"), ErrorKindUser)

    ErrHttpResourceNotFound = NewError(
        "errHttpResuorceNotFound", errors.New("resource not found"), ErrorKindUser)

   //ErrServiceResourceNotFound = NewError(
   //    "errServiceResourceNotFound", errors.New("resource is not found"), ErrorKindUser)
   //
   //ErrServiceResourceNotCreated = NewError(
   //    "errServiceResourceNotCreated", errors.New("resource is not created"), ErrorKindUser)
   //
   //ErrServiceResourceNotUpdated = NewError(
   //    "errServiceResourceNotUpdated", errors.New("resource is not updated"), ErrorKindUser)
   //
   //ErrServiceResourceNotRemoved = NewError(
   //    "errServiceResourceNotRemoved", errors.New("resource is not removed"), ErrorKindUser)
)
