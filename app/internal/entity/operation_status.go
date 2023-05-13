package entity

import (
    "encoding/json"
    "fmt"
)

type OperationStatus uint8

const (
    _ OperationStatus = iota
    OperationStatusConfirming
    OperationStatusUpdating
    OperationStatusConfirmed
)

var (
    OperationStatusName = map[OperationStatus]string{
        OperationStatusConfirming: "CONFIRMING",
        OperationStatusUpdating: "UPDATING",
        OperationStatusConfirmed: "CONFIRMED",
    }

    OperationStatusValue = map[string]OperationStatus{
        "CONFIRMING": OperationStatusConfirming,
        "UPDATING": OperationStatusUpdating,
        "CONFIRMED": OperationStatusConfirmed,
    }
)

func (as OperationStatus) String() string {
    return OperationStatusName[as]
}

func (as OperationStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(as.String())
}

func (as *OperationStatus) UnmarshalJSON(data []byte) error {
    var value string
    var err error

    if err = json.Unmarshal(data, &value); err != nil {
        return err
    }

    *as, err = parseOperationStatus(value)

    return err
}

func parseOperationStatus(value string) (OperationStatus, error) {
    if parsedValue, ok := OperationStatusValue[value]; ok {
        return parsedValue, nil
    }

    return OperationStatus(0), fmt.Errorf("%q is not a valid OperationStatus", value)
}
