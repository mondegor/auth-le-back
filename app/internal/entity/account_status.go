package entity

import (
    "encoding/json"
    "fmt"
)

type AccountStatus uint8

const (
    _ AccountStatus = iota
    AccountStatusActivating
    AccountStatusEnabled
    AccountStatusBlocked
    AccountStatusRemoved
)

var (
    AccountStatusName = map[AccountStatus]string{
        AccountStatusActivating: "ACTIVATING",
        AccountStatusEnabled: "ENABLED",
        AccountStatusBlocked: "BLOCKED",
        AccountStatusRemoved: "REMOVED",
    }

    AccountStatusValue = map[string]AccountStatus{
        "ACTIVATING": AccountStatusActivating,
        "ENABLED": AccountStatusEnabled,
        "BLOCKED": AccountStatusBlocked,
        "REMOVED": AccountStatusRemoved,
    }
)

func (as AccountStatus) String() string {
    return AccountStatusName[as]
}

func (as AccountStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(as.String())
}

func (as *AccountStatus) UnmarshalJSON(data []byte) error {
    var value string
    var err error

    if err = json.Unmarshal(data, &value); err != nil {
        return err
    }

    *as, err = parseAccountStatus(value)

    return err
}

func parseAccountStatus(value string) (AccountStatus, error) {
    if parsedValue, ok := AccountStatusValue[value]; ok {
        return parsedValue, nil
    }

    return AccountStatus(0), fmt.Errorf("%q is not a valid AccountStatus", value)
}
