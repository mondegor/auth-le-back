package postgresql

import "github.com/jackc/pgx/v5"

type queryRow struct {
    conn *Connection
    row pgx.Row
}

func wrapQueryRow(conn *Connection, row pgx.Row) pgx.Row {
    return &queryRow{
        conn: conn,
        row: row,
    }
}

func (qr *queryRow) Scan(dest ...any) error {
    err := qr.row.Scan(dest...)

    if err != nil {
        return qr.conn.wrapError(err)
    }

    return nil
}
