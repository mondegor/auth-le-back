package postgresql

import (
    "context"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgconn"
)

func (c *Connection) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
    c.debugQuery(sql)

    description, err := c.conn.Prepare(ctx, name, sql)

    if err != nil {
        return nil, c.wrapError(err)
    }

    return description, nil
}

func (c *Connection) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
    c.debugQuery(sql)

    commandTag, err := c.conn.Exec(ctx, sql, args...)

    if err != nil {
        return commandTag, c.wrapError(err)
    }

    return commandTag, nil
}

func (c *Connection) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
    c.debugQuery(sql)

	rows, err := c.conn.Query(ctx, sql, args...)

    if err != nil {
        return nil, c.wrapError(err)
    }

    return rows, nil
}

func (c *Connection) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
    c.debugQuery(sql)

    row := c.conn.QueryRow(ctx, sql, args...)

    return wrapQueryRow(c, row)
}
