package mrhttp

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "strconv"
)

type RequestPath struct {
    params httprouter.Params
    request *http.Request
}

func NewRequestPath(request *http.Request) RequestPath {
    ctx := request.Context()
    var v = ctx.Value(httprouter.ParamsKey)

    params, ok := v.(httprouter.Params)

    if !ok {
        params = nil
    }

    return RequestPath{
        params: params,
        request: request,
    }
}

func (rp RequestPath) Get(name string) string {
    if rp.params == nil {
        return ""
    }

    return rp.params.ByName(name)
}

func (rp RequestPath) GetInt(name string) int {
    value, err := strconv.Atoi(rp.Get(name))

    if err != nil {
        return 0
    }

    return value
}
