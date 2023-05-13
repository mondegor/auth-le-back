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
    Detail  string `json:"detail"`
    Request  string `json:"request"`
    Time  string `json:"time"`
    ErrorTraceId string `json:"errorTraceId"`
}

//// application/json
//type mrapp.ErrorAttribute struct {
//    Id string `json:"id"`
//    Value string `json:"value"`
//}

// application/json
type AppError400Response mrapp.ErrorList


func (ar *AppErrorResponse) Marshal() []byte {
    bytes, err := json.Marshal(ar)

    if err == nil {
        return bytes
    }

    return nil
}


func (ar *AppError400Response) AddError(id string, value string) {
   *ar = append(*ar, mrapp.ErrorAttribute{id, value})
}

func SendResponseError(w http.ResponseWriter, r *http.Request, err error) {
    w.Header().Set("Content-Type", "application/problem+json")

    var statusCode = http.StatusTeapot
    var title = "Internal server error"

    var appError *mrapp.AppError

    if errors.As(err, &appError) {

        if errors.Is(err, mrapp.ErrHttpResourceNotFound) {
            statusCode = http.StatusNotFound
            title = "Page not found"
        } else {
            if appError.Kind() == mrapp.ErrorKindUser {
                statusCode = http.StatusInternalServerError
            } else if appError.Kind() == mrapp.ErrorKindSystem {
                statusCode = http.StatusInternalServerError
            }
        }

        statusCode = http.StatusBadRequest
        //w.Write(appErr.Marshal())

        //w.Write(response.Marshal())
        return
    }

    // dd := mrapp.NewSystemError(err)
    w.WriteHeader(statusCode)

    wrappedResponse := AppErrorResponse{
        Title: title,
        Detail: err.Error(),
        Request: r.URL.Path,
        Time: time.Now().String(),
        ErrorTraceId: "fffffff",
    }

    // wrappedResponse.AddError("system", fmt.Sprintf("Code %s, error %s:", dd.Code(), dd.Error()))

    w.Write(wrappedResponse.Marshal())
}

func SendResponseError400(w http.ResponseWriter, errors *mrapp.ErrorList) error {
    w.WriteHeader(http.StatusBadRequest)

    if len(*errors) == 0 {
        w.Write([]byte("[]"))
        return nil
    }

    bytes, err := json.Marshal(*errors)

    if err != nil {
        return mrapp.ErrHttpResponseParseData.Wrap(err)
    }

    w.Write(bytes)

    return nil
}

func SendResponseSuccessNoBody(w http.ResponseWriter) error {
    w.WriteHeader(http.StatusNoContent)
    return nil
}

//func systemError(err error) *AppError {
//    return NewAppError("US-000000", "internal system error", err.Error(), err)
//}

func SendResponse(w http.ResponseWriter, status int, response any) {
    w.WriteHeader(status)

    //wrappedResponse := AppResponse{
    //    TimeStamp: time.Now(),
    //    StatusCode: status,
    //    Response: response,
    //}

    // w.Write(response.Marshal())
}

func SendResponseNoContent(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNoContent)
}
