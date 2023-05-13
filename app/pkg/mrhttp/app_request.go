package mrhttp

import (
    "auth-le-back/pkg/mrapp"
    "encoding/json"
    "net/http"
)

func BindJSON(r *http.Request, v any) error {
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    if err := dec.Decode(&v); err != nil {
        return mrapp.ErrHttpRequestParseData.Wrap(err)
    }

    return nil
}

