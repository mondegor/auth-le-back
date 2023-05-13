package entity

import (
    "encoding/json"
    "fmt"
)

type UserStatus uint8

const (
    _ UserStatus = iota
    UserStatusActivating
    UserStatusEnabled
    UserStatusBlocked
    UserStatusRemoved
)

var (
    UserStatusName = map[UserStatus]string{
        UserStatusActivating: "ACTIVATING",
        UserStatusEnabled: "ENABLED",
        UserStatusBlocked: "BLOCKED",
        UserStatusRemoved: "REMOVED",
    }

    UserStatusValue = map[string]UserStatus{
        "ACTIVATING": UserStatusActivating,
        "ENABLED": UserStatusEnabled,
        "BLOCKED": UserStatusBlocked,
        "REMOVED": UserStatusRemoved,
    }
)

func (us UserStatus) String() string {
    return UserStatusName[us]
}

func (us UserStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(us.String())
}

func (us *UserStatus) UnmarshalJSON(data []byte) error {
    var value string
    var err error

    if err = json.Unmarshal(data, &value); err != nil {
        return err
    }

    *us, err = parseUserStatus(value)

    return err
}

func parseUserStatus(value string) (UserStatus, error) {
    if parsedValue, ok := UserStatusValue[value]; ok {
        return parsedValue, nil
    }

    return UserStatus(0), fmt.Errorf("%q is not a valid UserStatus", value)
}
