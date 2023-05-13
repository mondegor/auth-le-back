package mrapp

var (
    ErrServiceResourceTemporarilyUnavailable = NewError(
        "errServiceResourceTemporarilyUnavailable", ErrorKindSystem, "resource is temporarily unavailable")

    ErrServiceResourceNotFound = NewError(
        "errServiceResourceNotFound", ErrorKindUser, "resource is not found")

    ErrServiceResourceNotCreated = NewError( // arg1=resourceName
        "errServiceResourceNotCreated", ErrorKindSystem, "resource is not created")

    ErrServiceResourceNotUpdated = NewError(
        "errServiceResourceNotUpdated", ErrorKindSystem, "resource is not updated")

    ErrServiceResourceNotRemoved = NewError(
        "errServiceResourceNotRemoved", ErrorKindSystem, "resource is not removed")

    ErrServiceIncorrectData = NewError(
        "errServiceIncorrectData", ErrorKindSystem, "data is incorrect")
)
