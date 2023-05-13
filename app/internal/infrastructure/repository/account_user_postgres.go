package repository

import (
    "auth-le-back/internal/entity"
    "auth-le-back/pkg/mrlib"
    "context"
)

func (a *Account) GetAccountInfo(ctx context.Context, id entity.AccountPrimaryKey) (*entity.AccountUser, error) {
	//sql := `
    //    SELECT
    //        aa.account_id, au.user_login, au.user_email,
    //        au.datetime_last_login, au.user_last_login_ip, aa.account_status
    //    FROM
    //        public.auth_accounts aa
    //    JOIN
    //        public.auth_users au
    //    ON
    //        aa.account_id = au.account_id
    //    WHERE aa.account_id = $1 AND aa.account_status <> $2;`
    //
	//item := &entity.AccountUser{}
    //
    //err := a.QueryRow(ctx, sql, id, entity.AccountStatusRemoved).Scan(
    //    &item.Account.Id,
    //    &item.User.Login,
    //    &item.User.Email,
    //    &item.User.LoggedAt,
    //    &item.User.LoggedIPInt,
    //    &item.Account.Status)

    var err error
    item := &entity.AccountUser{}

    item.Account.Id = "token-id"
    item.User.Login = "user-login"
    item.User.Email = "user@login"
    item.User.LoggedAt = "2020-01-23"
    item.User.LoggedIPInt = 37747432
    item.Account.Status = entity.AccountStatusEnabled

    item.User.LoggedIP = mrlib.Int2ip(item.User.LoggedIPInt)

	return item, err
}
