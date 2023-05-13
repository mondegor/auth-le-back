package http_v1

import (
    "auth-le-back/internal/controller/dto"
    "auth-le-back/internal/entity"
    "auth-le-back/internal/usecase"
    "auth-le-back/pkg/mrapp"
    "auth-le-back/pkg/mrhttp"
    "encoding/json"
    "net/http"
)

const (
    authSignupURL = "/v1/signup"
    authSigninURL = "/v1/signin"
    authURL       = "/v1/auth"
    authTokenURL  = "/v1/auth:token"
)

type Auth struct {
    logger mrapp.Logger
    validator mrapp.Validator
    auth usecase.AuthService
}

func NewAuth(logger mrapp.Logger, validator mrapp.Validator, auth usecase.AuthService) *Auth {
    return &Auth{
        logger: logger,
        validator: validator,
        auth: auth,
    }
}

func (a *Auth) AddHandlers(router mrapp.Router) {
    router.HandlerFunc(http.MethodPost, authSignupURL, a.Signup())
    router.HandlerFunc(http.MethodPost, authSigninURL, a.Signin)
    router.HandlerFunc(http.MethodGet, authURL, a.GetAccount)
    router.HandlerFunc(http.MethodPost, authURL, a.GenerateTokens)
    router.HandlerFunc(http.MethodGet, authTokenURL, a.GenerateTokensBy)
}

func (a *Auth) Signup() mrapp.HttpHandlerFunc {
    type CreateAccount struct {
        UserEmail string `json:"userEmail" validate:"required,min=7,max=128,email"`
    }

    return func(w http.ResponseWriter, r *http.Request) error {
        request := CreateAccount{}

        if err := mrhttp.BindJSON(r, &request); err != nil {
            return err
        }

        if err := a.validator.Validate(request); err != nil {
            return err
        }

        accountId, err := a.auth.CreateAccount(
            r.Context(),
            entity.AccountUser{
                User: entity.User{Email: request.UserEmail},
            })

        if err != nil {
            return err
        }

        // todo operationToken, accountId, message
        return mrhttp.SendResponse(w, http.StatusCreated, accountId)
    }
}

func (a *Auth) Signin(w http.ResponseWriter, r *http.Request) error {

    response := &mrapp.ErrorList{}
    response.Add("ff", "dddddd")
    response.Add("ff1", "dddddd2")

    w.WriteHeader(http.StatusBadRequest)

    bytes, err := json.Marshal(response)

    if err != nil {
        return err
    }

    w.WriteHeader(http.StatusOK)
    w.Write(bytes)

    return nil
}

func (a *Auth) GetAccount(w http.ResponseWriter, r *http.Request) error {

    var response dto.AccountResponse

    account, err := a.auth.GetAccount(r.Context(), "acc-id")

    if err != nil {
        return err
    }

    response.Init(account)

    bytes, err := json.Marshal(response)

    if err != nil {
        return err
    }

    // mrhttp.SendResponse(w, http.StatusOK, response)
    // h.logger.Info("This is Get")

    w.WriteHeader(http.StatusOK)
    w.Write(bytes)

    return nil
}

func (a *Auth) GenerateTokens(w http.ResponseWriter, r *http.Request) error {

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("GenerateTokens"))

    return nil
}

func (a *Auth) GenerateTokensBy(w http.ResponseWriter, r *http.Request) error {

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("GenerateTokensBy"))

    return nil
}
