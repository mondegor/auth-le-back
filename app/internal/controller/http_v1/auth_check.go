package http_v1

import (
    "auth-le-back/internal/controller/dto"
    "auth-le-back/internal/usecase"
    "auth-le-back/pkg/mrapp"
    "auth-le-back/pkg/mrhttp"
    "net/http"
)

const (
    authCheckLoginCheck = "/v1/auth/login-check"
    authCheckEmailCheck = "/v1/auth/email-check"
)

type AuthCheck struct {
    logger mrapp.Logger
    validator mrapp.Validator
    auth usecase.AuthService
}

func NewAuthCheck(logger mrapp.Logger, validator mrapp.Validator, auth usecase.AuthService) *AuthCheck {
    return &AuthCheck{
        logger: logger,
        validator: validator,
        auth: auth,
    }
}

func (a *AuthCheck) AddHandlers(router mrapp.Router) {
    router.HandlerFunc(http.MethodPost, authCheckLoginCheck, a.LoginCheck())
    router.HandlerFunc(http.MethodPost, authCheckEmailCheck, a.EmailCheck())
}

func (a *AuthCheck) LoginCheck() mrapp.HttpHandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) error {
        request := dto.CheckLoginValue{}

        if err := mrhttp.BindJSON(r, &request); err != nil {
            return err
        }
        if err := a.validator.Validate(request); err != nil {
            return err
        }

        err := a.auth.CheckIfLoginIsFree(r.Context(), request.UserLogin)

        if err != nil {
            return err
        }

        return mrhttp.SendResponseNoContent(w)
    }
}

func (a *AuthCheck) EmailCheck() mrapp.HttpHandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) error {
        request := dto.CheckEmailValue{}

        if err := mrhttp.BindJSON(r, &request); err != nil {
            return err
        }
        if err := a.validator.Validate(request); err != nil {
            return err
        }

        err := a.auth.CheckIfEmailIsFree(r.Context(), request.UserEmail)

        if err != nil {
            return err
        }

        return mrhttp.SendResponseNoContent(w)
    }
}
