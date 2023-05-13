package repository

import (
    "auth-le-back/internal/entity"
    "auth-le-back/pkg/client/postgresql"
    "auth-le-back/pkg/mrapp"
    "context"
)

type User struct {
    *postgresql.Connection
}

func NewUser(client *postgresql.Connection) *User {
	return &User{client}
}

func (a *User) FindOne(ctx context.Context, id entity.UserPrimaryKey) (*entity.User, error) {
	sql := `
        SELECT
            user_id, account_id, visitor_id, user_type, datetime_created,
            user_login, user_email,
            datetime_last_login, user_last_login_ip, datetime_last_visit,
            user_status, datetime_status
        FROM
            public.auth_users
        WHERE user_id = $1 AND user_status <> $2;`

	item := &entity.User{}

    err := a.QueryRow(ctx, sql, id, entity.UserStatusRemoved).Scan(
        &item.Id,
        &item.CreatedAt,
        &item.Status,
        &item.ChangedAt)

	return item, err
}

func (a *User) IsExists(ctx context.Context, id entity.UserPrimaryKey) (bool, error) {
	sql := `
        SELECT 1
        FROM
            public.auth_users
        WHERE user_id = $1 AND user_status <> $2;`

	var isExists bool

    err := a.QueryRow(ctx, sql, id, entity.UserStatusRemoved).Scan(&isExists)

	return isExists, err
}

func (a *User) Create(ctx context.Context, row *entity.User) error {
	sql := `
        INSERT INTO public.auth_users
            (account_id, visitor_id, user_type, datetime_created,
             user_login, user_email,
             user_status, datetime_status)
        VALUES
            ($1, $2, $3, NOW(), $4, $5, $6, NOW())
        RETURNING user_id, datetime_created, datetime_status;`

	err := a.QueryRow(
	    ctx,
		sql,
		row.AccountId,
        row.VisitorId,
        row.UserType,
        row.Login,
        row.Email,
		row.Status,
        entity.UserStatusActivating).Scan(
            &row.Id,
            &row.CreatedAt,
            &row.ChangedAt)

    return err
}

func (a *User) Update(ctx context.Context, row *entity.User) error {
	sql := `
        UPDATE public.auth_users
        SET user_status = $2,
            datetime_status = NOW()
        WHERE user_id = $1;`

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

func (a *User) Delete(ctx context.Context, id entity.UserPrimaryKey) error {
	sql := `
        UPDATE public.auth_users
        SET
            user_status = $2
        WHERE
            user_id = $1`

	commandTag, err := a.Exec(ctx, sql, id, entity.UserStatusRemoved)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return mrapp.ErrStorageNoRowFound
	}

	return nil
}

func (a *User) GetIdByLogin(ctx context.Context, login string) (entity.UserPrimaryKey, error) {
    sql := `
        SELECT user_id
        FROM
            public.auth_users
        WHERE user_login = $1;`

    var userId entity.UserPrimaryKey

    err := a.QueryRow(ctx, sql, login).Scan(&userId)

    return userId, err
}

func (a *User) GetIdByEmail(ctx context.Context, email string) (entity.UserPrimaryKey, error) {
    sql := `
        SELECT user_id
        FROM
            public.auth_users
        WHERE user_email = $1;`

    var userId entity.UserPrimaryKey

    err := a.QueryRow(ctx, sql, email).Scan(&userId)

    return userId, err
}
