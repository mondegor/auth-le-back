package repository

import (
    "auth-le-back/internal/entity"
    "auth-le-back/pkg/client/postgresql"
    "auth-le-back/pkg/mrapp"
    "context"
)

type Account struct {
    *postgresql.Connection
}

func NewAccount(client *postgresql.Connection) *Account {
	return &Account{client}
}

func (a *Account) FindOne(ctx context.Context, id entity.AccountPrimaryKey) (*entity.Account, error) {
	sql := `
        SELECT
            account_id, datetime_created, account_status, datetime_status
        FROM
            public.auth_accounts
        WHERE account_id = $1 AND account_status <> $2;`

	item := &entity.Account{}

    err := a.QueryRow(ctx, sql, id, entity.AccountStatusRemoved).Scan(
        &item.Id,
        &item.CreatedAt,
        &item.Status,
        &item.ChangedAt)

	return item, err
}

func (a *Account) IsExists(ctx context.Context, id entity.AccountPrimaryKey) (bool, error) {
	sql := `
        SELECT 1
        FROM
            public.auth_accounts
        WHERE account_id = $1 AND account_status <> $2;`

	var isExists bool

    err := a.QueryRow(ctx, sql, id, entity.AccountStatusRemoved).Scan(&isExists)

	return isExists, err
}

func (a *Account) Create(ctx context.Context, row *entity.AccountUser) error {
	sql := `
        INSERT INTO public.auth_accounts
            (datetime_created, account_status, datetime_status)
        VALUES
            (NOW(), $1, NOW())
        RETURNING account_id, datetime_created;`

	err := a.QueryRow(
	    ctx,
		sql,
        entity.AccountStatusActivating).Scan(
            &row.Account.Id,
            &row.Account.CreatedAt)

    return err
}

func (a *Account) Update(ctx context.Context, row *entity.Account) error {
	sql := `
        UPDATE public.auth_accounts
        SET account_status = $2,
            datetime_status = NOW()
        WHERE account_id = $1;`

	commandTag, err := a.Exec(
	    ctx,
	    sql,
	    row.Id,
	    row.Status)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return mrapp.ErrStorageNoRowFound
	}

	return nil
}

func (a *Account) Delete(ctx context.Context, id entity.AccountPrimaryKey) error {
	sql := `
        UPDATE public.auth_accounts
        SET
            account_status = $2
        WHERE
            account_id = $1`

	commandTag, err := a.Exec(ctx, sql, id, entity.AccountStatusRemoved)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return mrapp.ErrStorageNoRowFound
	}

	return nil
}
