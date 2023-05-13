package mrhttp

import (
    "auth-le-back/pkg/mrapp"
    "encoding/json"
    "errors"
    "net/http"
    "time"
)

//ErrBadRequest       = RegisterErrorType(BaseErr, http.StatusBadRequest, ErrCodeBadRequest)             // 400
//ErrUnauthorized     = RegisterErrorType(BaseErr, http.StatusUnauthorized, ErrCodeUnauthorized)         // 401
//ErrPaymentRequired  = RegisterErrorType(BaseErr, http.StatusPaymentRequired, ErrCodePaymentRequired)   // 402
//ErrForbidden        = RegisterErrorType(BaseErr, http.StatusForbidden, ErrCodeForbidden)               // 403
//ErrNotFound         = RegisterErrorType(BaseErr, http.StatusNotFound, ErrCodeNotFound)                 // 404
//ErrMethodNotAllowed = RegisterErrorType(BaseErr, http.StatusMethodNotAllowed, ErrCodeMethodNotAllowed) // 405

// application/json
type AppSuccessResponse struct {
    Success bool `json:"success"`
}

// application/problem+json:
type AppErrorResponse struct {
    Title string `json:"title"`
    Details  string `json:"details"`
    Request  string `json:"request"`
    Time  string `json:"time"`
    ErrorTraceId string `json:"errorTraceId"`
}

//// application/json
//type mrapp.ErrorAttribute struct {
//    Id string `json:"id"`
//    Value string `json:"value"`
//}

func (ar *AppErrorResponse) Marshal() []byte {
    bytes, err := json.Marshal(ar)

    if err == nil {
        return bytes
    }

    return nil
}

func SendResponseError(w http.ResponseWriter, r *http.Request, headers *RequestHeaders, err error) {
    if errorList, ok := err.(*mrapp.ErrorList); ok {
        if err = SendResponse(w, http.StatusBadRequest, errorList); err == nil {
            return
        }
    }

    w.Header().Set("Content-Type", "application/problem+json")

    appError, ok := err.(*mrapp.AppError)

    if !ok {
        appError = mrapp.ErrInternal.Wrap(err).(*mrapp.AppError)
    }

    var statusCode = http.StatusTeapot

    errMessage := appError.UserError(headers.Locale)

    if errors.Is(err, mrapp.ErrHttpResourceNotFound) {
        statusCode = http.StatusNotFound
    } else {
        statusCode = http.StatusInternalServerError
    }

    w.WriteHeader(statusCode)

    errorResponse := AppErrorResponse{
        Title: errMessage.Reason,
        Details: errMessage.GetDetails(),
        Request: r.URL.Path,
        Time: time.Now().Format(time.RFC3339),
        ErrorTraceId: headers.CorrelationId,
    }

    w.Write(errorResponse.Marshal())
}

func SendResponseNoContent(w http.ResponseWriter) error {
    w.WriteHeader(http.StatusNoContent)

    return nil
}

func SendResponse(w http.ResponseWriter, status int, response any) error {
    w.WriteHeader(status)

    bytes, err := json.Marshal(response)

    if err != nil {
        return mrapp.ErrHttpResponseParseData.Wrap(err)
    }

    _, err = w.Write(bytes)

    if err != nil {
        return mrapp.ErrHttpResponseSendData.Wrap(err)
    }

    return nil
}
