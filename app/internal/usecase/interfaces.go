package usecase

import (
    "auth-le-back/internal/entity"
    "context"
)

type (
    AuthService interface {
        CreateAccount(ctx context.Context, item entity.AccountUser) (entity.AccountPrimaryKey, error)
        AuthAccount(ctx context.Context, itemId entity.AccountPrimaryKey) (*entity.Account, error)
        GetAccount(ctx context.Context, itemId entity.AccountPrimaryKey) (*entity.AccountUser, error)
        GenerateTokens(ctx context.Context, item *entity.Account) error
        GenerateTokensBy(ctx context.Context, item *entity.Account) error

        // AuthChangeService
        AuthCheckService
    }

    AuthChangeService interface {
        ChangeLogin(ctx context.Context, item *entity.User) (entity.AccountPrimaryKey, error)
        ChangeEmail(ctx context.Context, itemId entity.AccountPrimaryKey) (*entity.Account, error)
    }

    AuthCheckService interface {
        IsLoginAlreadyExist(ctx context.Context, login string) error
    }

    AccountStorage interface {
       FindOne(ctx context.Context, id entity.AccountPrimaryKey) (*entity.Account, error)
       IsExists(ctx context.Context, id entity.AccountPrimaryKey) (bool, error)
       Create(ctx context.Context, row *entity.AccountUser) error
       Update(ctx context.Context, row *entity.Account) error
       Delete(ctx context.Context, id entity.AccountPrimaryKey) error
       GetAccountInfo(ctx context.Context, id entity.AccountPrimaryKey) (*entity.AccountUser, error)
    }

    UserStorage interface {
        FindOne(ctx context.Context, id entity.UserPrimaryKey) (*entity.User, error)
        Create(ctx context.Context, row *entity.User) error
        Update(ctx context.Context, row *entity.User) error
        Delete(ctx context.Context, id entity.UserPrimaryKey) error
        GetIdByLogin(ctx context.Context, login string) (entity.UserPrimaryKey, error)
        GetIdByEmail(ctx context.Context, email string) (entity.UserPrimaryKey, error)
    }
)
