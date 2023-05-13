package usecase

import (
    "context"
    "fmt"
    "net/mail"
    "regexp"
)

func (a *Auth) IsLoginAlreadyExist(ctx context.Context, login string) error {
    if _, err := mail.ParseAddress(login); err == nil {
        return a.user.IsEmailExists(ctx, login)
    }



    if ok, _ := regexp.MatchString("^[a-z][a-z0-9]*$", login); ok == true {
        return a.user.IsLoginExists(ctx, login)
    }

    return true, fmt.Errorf("Login is incorrect")
}
