package mrhttp

import (
    "auth-le-back/pkg/mrapp"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

// go get -u github.com/julienschmidt/httprouter

// Make sure the Router conforms with the mrapp.Router interface
var _ mrapp.Router = (*Router)(nil)

type Router struct {
    logger mrapp.Logger
    router *httprouter.Router
    generalHandler http.Handler
}

func NewRouter(logger mrapp.Logger) *Router {
    router := httprouter.New()

    // rt.router.NotFound
    // rt.router.MethodNotAllowed

    return &Router{
        logger: logger,
        router: router,
        generalHandler: router,
    }
}

func (rt *Router) RegisterMiddleware(handlers ...mrapp.HttpMiddleware) {
    // recursion call: handler1(handler2(handler3(router())))
    for i := len(handlers) - 1; i >= 0; i-- {
        rt.generalHandler = handlers[i].Middleware(rt.generalHandler)
    }
}

func (rt *Router) Register(controllers ...mrapp.HttpController) {
    for _, controller := range controllers {
        controller.AddHandlers(rt)
    }
}

func (rt *Router) HandlerFunc(method, path string, handler mrapp.HttpHandlerFunc) {
    rt.router.Handler(method, path, rt.MiddlewareLast(handler))
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    rt.generalHandler.ServeHTTP(w, r)
}
