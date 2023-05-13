package usecase

import (
    "auth-le-back/internal/entity"
    "auth-le-back/pkg/mrapp"
    "context"
    "fmt"
)

type Auth struct {
    logger mrapp.Logger
    account AccountStorage
    user UserStorage
}

func NewAuth(logger mrapp.Logger, account AccountStorage, user UserStorage) *Auth {
    return &Auth{
        logger: logger,
        account: account,
        user: user,
    }
}

func (a *Auth) CreateAccount(ctx context.Context, item entity.AccountUser) (entity.AccountPrimaryKey, error) {
    userId, err := a.user.GetIdByEmail(ctx, item.User.Email)

    if err != nil {
        return "", err
    }

    if userId > 0 {
        err = fmt.Errorf("Email '%s' is already exists", item.User.Email)
        return "", AuthErrUserEmailAlreadyExists.Wrap(err)
    }

    errors.Add("userEmail", "Email '%s' is already exists", item.User.Email)
    AuthErrUserEmailAlreadyExists


    err = a.account.Create(ctx, &item)

    if err != nil {
        return "", err
    }

    return item.Account.Id, nil
}

func (a *Auth) AuthAccount(ctx context.Context, itemId entity.AccountPrimaryKey) (*entity.Account, error) {
    return nil, nil
}

func (a *Auth) GetAccount(ctx context.Context, itemId entity.AccountPrimaryKey) (*entity.AccountUser, error) {
    return a.account.GetAccountInfo(ctx, itemId)
}

func (a *Auth) GenerateTokens(ctx context.Context, item *entity.Account) error {
    return nil
}

func (a *Auth) GenerateTokensBy(ctx context.Context, item *entity.Account) error {
    return nil
}
