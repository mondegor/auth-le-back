package entity

type (
    OperationToken string

    // Operation - операция
    Operation struct { // DB: auth_operation
        Id            OperationToken // operation_token
        maxCheckCode  AccountPrimaryKey // max_check_code
        sessionExpiry string // session_expiry
        countTryCode  int32 // count_try_code
        Status        OperationStatus // operation_status
    }
)
