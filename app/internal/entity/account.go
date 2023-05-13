package entity

import "time"

type (
    AccountPrimaryKey string

    // Account - учётная запись пользователя
    Account struct { // DB: auth_accounts
        Id        AccountPrimaryKey // account_id
        CreatedAt time.Time     // datetime_created
        Status    AccountStatus // account_status
        ChangedAt string        // datetime_status
    }

    // AccountUser - учётная запись пользователя
    AccountUser struct {
        Account Account
        User User
    }
)
