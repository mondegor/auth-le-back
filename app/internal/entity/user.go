package entity

import "net"

type (
    UserPrimaryKey int

    // User - пользователя
    User struct { // DB: auth_user
        Id          UserPrimaryKey // user_id
        AccountId   AccountPrimaryKey // account_id
        VisitorId   string     // visitor_id
        UserType    UserType   // user_type
        CreatedAt   string     // datetime_created
        Login       string     // user_login
        Email       string     // user_email
        LoggedAt    string     // datetime_last_login
        LoggedIPInt uint32     // user_last_login_ip
        LoggedIP    net.IP
        VisitedAt   string     // datetime_last_visit
        Status      UserStatus // user_status
        ChangedAt   string     // datetime_status
    }
)
