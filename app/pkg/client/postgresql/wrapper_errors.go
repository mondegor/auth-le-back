package postgresql

import (
    "auth-le-back/pkg/mrapp"
    "errors"

    "github.com/jackc/pgx/v5/pgconn"
)

func (c *Connection) wrapError(err error) error {
    var pgErr *pgconn.PgError

    if errors.As(err, &pgErr) {
        // Severity: ERROR; Code: 42601; Message syntax error at or near "field_status"
        return mrapp.ErrStorageQueryFailed.Wrap(err)
    }

    if err.Error() == "no rows in result set" {
        return mrapp.ErrStorageNoRowFound.Wrap(err)
    }

    return err
}
