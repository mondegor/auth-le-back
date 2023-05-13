package usecase

import (
    "auth-le-back/internal/entity"
    "auth-le-back/pkg/mrapp"
    "context"
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
        return "", mrapp.ErrServiceResourceTemporarilyUnavailable.Wrap(err)
    }

    if userId > 0 {
        return "", ErrAuthUserEmailAlreadyExists.New(item.User.Email)
    }

    item.Account.Status = entity.AccountStatusActivating

    a.logger.Info("account.Create: %s", item.User.Email)

    err = a.account.Create(ctx, &item)

    if err != nil {
        return "", mrapp.ErrServiceResourceNotCreated.Wrap(err)
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
