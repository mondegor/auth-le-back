package mrhttp

import (
    "auth-le-back/pkg/mrapp"
    "context"
    "net/http"
)

const (
    ctxKeyHeadersID ctxKey = iota
    PlatformWeb = "WEB"
    PlatformMobile = "MOBILE"
)

type ctxKey uint8

type RequestHeaders struct {
    Locale mrapp.Locale
    CorrelationId string
    Platform string
}

func (rt *Router) MiddlewareLast(h mrapp.HttpHandlerFunc) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        rh, _ := r.Context().Value(ctxKeyHeadersID).(RequestHeaders)
        rt.logger.Info("MESSAGE-language: %s", rh.Locale.GetCode())
        rt.logger.Info("MESSAGE-correlationId: %s", rh.CorrelationId)

        err := h(w, r)

        if err != nil {
            rt.logger.Error(err)

            SendResponseError(w, r, err)
        }
    })
}

func MiddlewareHeaders(logger mrapp.Logger, translator mrapp.Translator) mrapp.HttpMiddleware {
    return mrapp.HttpMiddlewareFunc(func(next http.Handler) http.Handler {
        logger.Info("Added MiddlewareHeaders")

        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            logger.Debug("Exec MiddlewareHeaders")

            rh := RequestHeaders{Platform: PlatformWeb}

            if r.Header.Get("Accept-Language") != "" {
                rh.Locale = translator.GetLocaleByAcceptLanguage(r.Header.Get("Accept-Language"))
            }

            if r.Header.Get("CorrelationID") != "" {
                rh.CorrelationId = r.Header.Get("CorrelationID")
            }

            if r.Header.Get("Platform") == PlatformMobile {
                rh.Platform = PlatformMobile
            }

            next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyHeadersID, rh)))
        })
    })
}

func MiddlewareAuthenticateUser(logger mrapp.Logger) mrapp.HttpMiddleware {
    return mrapp.HttpMiddlewareFunc(func(next http.Handler) http.Handler {
        logger.Info("Added MiddlewareAuthenticateUser")

        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            logger.Debug("Exec MiddlewareAuthenticateUser")

            //if len(p) > 0 {
            //    ctx := req.Context()
            //    ctx = context.WithValue(ctx, ParamsKey, p)
            //    req = req.WithContext(ctx)
            //}

            next.ServeHTTP(w, r)
        })
    })
}
