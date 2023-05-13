package usecase

import (
    "auth-le-back/pkg/mrapp"
    "context"
)

func (a *Auth) CheckIfLoginIsFree(ctx context.Context, login string) error {
    userId, err := a.user.GetIdByEmail(ctx, login)

    if err != nil {
        return mrapp.ErrServiceResourceTemporarilyUnavailable.Wrap(err)
    }

    if userId > 0 {
        return ErrAuthUserLoginAlreadyExists.New(login)
    }

    return nil
}

func (a *Auth) CheckIfEmailIsFree(ctx context.Context, email string) error {
    userId, err := a.user.GetIdByEmail(ctx, email)

    if err != nil {
        return mrapp.ErrServiceResourceTemporarilyUnavailable.Wrap(err)
    }

    if userId > 0 {
        return ErrAuthUserEmailAlreadyExists.New(email)
    }

    return nil
}
