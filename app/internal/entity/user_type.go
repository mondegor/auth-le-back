package entity

import (
    "encoding/json"
    "fmt"
)

type UserType uint8

const (
    _ UserType = iota
    UserTypeSystem
    UserTypeUser
)

var (
    UserTypeName = map[UserType]string{
        UserTypeSystem: "SYSTEM",
        UserTypeUser: "USER",
    }

    UserTypeValue = map[string]UserType{
        "SYSTEM": UserTypeSystem,
        "USER": UserTypeUser,
    }
)

func (ut UserType) String() string {
    return UserTypeName[ut]
}

func (ut UserType) MarshalJSON() ([]byte, error) {
    return json.Marshal(ut.String())
}

func (ut *UserType) UnmarshalJSON(data []byte) error {
    var value string
    var err error

    if err = json.Unmarshal(data, &value); err != nil {
        return err
    }

    *ut, err = parseUserType(value)

    return err
}

func parseUserType(value string) (UserType, error) {
    if parsedValue, ok := UserTypeValue[value]; ok {
        return parsedValue, nil
    }

    return UserType(0), fmt.Errorf("%q is not a valid UserType", value)
}
