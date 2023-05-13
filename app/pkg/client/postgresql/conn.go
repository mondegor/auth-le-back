package postgresql

import (
    "auth-le-back/pkg/mrapp"
    "context"
    "fmt"
    "time"

    "github.com/jackc/pgx/v5"
)

// go get -u github.com/jackc/pgx/v5

type (
	Connection struct {
        conn *pgx.Conn
        logger mrapp.Logger
    }

    Options struct {
        Host string
        Port string
        Database string
        Username string
        Password string
        MaxPoolSize int32
        ConnAttempts int32
        ConnTimeout time.Duration
    }
)

func New(logger mrapp.Logger) *Connection {
	return &Connection{
        logger: logger,
    }
}

func (c *Connection) Connect(ctx context.Context, opt Options) error {
	if c.conn != nil {
		return mrapp.ErrStorageConnectionAlreadyExists
	}

	ctx, cancel := context.WithTimeout(ctx, opt.ConnTimeout * time.Second)
	defer cancel()

	var err error
	c.conn, err = pgx.Connect(ctx, getConnString(&opt))

	if err != nil {
        return mrapp.ErrStorageConnectionFailed.Wrap(err)
    }

    return nil
}

func (c *Connection) Close(ctx context.Context) error {
	if c.conn == nil {
		panic("connection had not opened")
	}

	conn := c.conn
	c.conn = nil

    if err := conn.Close(ctx); err != nil {
        return mrapp.ErrStorageConnectionFailed.Wrap(err)
    }

	return nil
}

func getConnString(o *Options) string {
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        o.Username,
        o.Password,
        o.Host,
        o.Port,
        o.Database)
}
