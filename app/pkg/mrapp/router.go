package mrapp

import "net/http"

type Router interface {
    RegisterMiddleware(handlers ...HttpMiddleware)
    Register(controllers ...HttpController)
    HandlerFunc(method, path string, handler HttpHandlerFunc)
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HttpMiddleware interface {
    Middleware(next http.Handler) http.Handler
}

type HttpMiddlewareFunc func(next http.Handler) http.Handler

func (f HttpMiddlewareFunc) Middleware(next http.Handler) http.Handler {
    return f(next)
}

type HttpController interface {
    AddHandlers(router Router)
}

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request) error
