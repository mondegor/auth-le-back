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

func (rt *Router) MiddlewareLast(h mrapp.HttpHandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var err error
        w.Header().Set("Content-Type", "application/json")

        rh, ok := r.Context().Value(ctxKeyHeadersID).(RequestHeaders)

        if ok {
            err = h(w, r)
        } else {
            err = mrapp.ErrInternalTypeAssertion
        }

        if err != nil {
            rt.logger.Info("MiddlewareLast:ERROR")
            rt.logger.Error(err)

            SendResponseError(w, r, &rh, err)
        }
    }
}

func MiddlewareHeaders(logger mrapp.Logger, translator mrapp.Translator) mrapp.HttpMiddleware {
    return mrapp.HttpMiddlewareFunc(func(next http.Handler) http.Handler {
        logger.Info("Added MiddlewareHeaders")

        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            logger.Debug("Exec MiddlewareHeaders")

            rh := RequestHeaders{Platform: PlatformWeb}

            acceptLanguage := r.Header.Get("Accept-Language")
            rh.Locale = translator.GetLocaleByAcceptLanguage(acceptLanguage)
            logger.Debug("Accept-Language: %s; Set-Language: %s", acceptLanguage, rh.Locale.GetCode())

            if r.Header.Get("CorrelationID") != "" {
                rh.CorrelationId = r.Header.Get("CorrelationID")
                logger.Debug("CorrelationID: %s", rh.CorrelationId)
            }

            if r.Header.Get("Platform") == PlatformMobile {
                rh.Platform = PlatformMobile
            }

            logger.Debug("Platform: %s", rh.Platform)

            next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyHeadersID, rh)))
        })
    })
}

func MiddlewareAuthenticateUser(logger mrapp.Logger) mrapp.HttpMiddleware {
    return mrapp.HttpMiddlewareFunc(func(next http.Handler) http.Handler {
        logger.Info("Added MiddlewareAuthenticateUser")

        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            logger.Debug("Exec MiddlewareAuthenticateUser")

            next.ServeHTTP(w, r)
        })
    })
}
