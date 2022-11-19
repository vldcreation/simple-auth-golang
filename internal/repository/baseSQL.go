package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/vldcreation/simple-auth-golang/internal/entity"
)

type (
	BeginTx interface {
		BeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sql.Tx, err error)
	}
	ExecContext interface {
		ExecContext(ctx context.Context, query string, args ...interface{}) (res sql.Result, err error)
	}
	PingContext interface {
		PingContext(ctx context.Context) (err error)
	}
	PrepareContext interface {
		PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error)
	}
	QueryContext interface {
		QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error)
	}
	QueryRowContext interface {
		QueryRowContext(ctx context.Context, query string, args ...interface{}) (row *sql.Row)
	}
)

// SQLConn is a common interface of *sql.DB and *sql.Conn.
type SQLConn interface {
	BeginTx
	io.Closer
	PingContext
	SQLTxConn
}

// SQLTxConn is a common interface of *sql.DB, *sql.Conn, and *sql.Tx.
type SQLTxConn interface {
	ExecContext
	PrepareContext
	QueryContext
	QueryRowContext
}

type SQL struct{}

var (
	ErrInvalidArgumentsScan = errors.New("Invalid arguments for scan")
	ErrNoColumnsReturned    = errors.New("No columns returned")
)

type boxExec struct {
	sqlResult sql.Result
	err       error
}

// WithDSN will open connection from the given dsn string with URL format, note
// that any error when opening the database should result in a panic.
func (SQL) WithDSN(dsn string) (conn *sql.DB) {
	driverName := strings.Split(dsn+"://", "://")[0]
	switch driverName {
	case "postgres", "postgresql":
		conn = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		if conn == nil {
			panic("empty database")
		}
	default:
		var err error

		conn, err = sql.Open(driverName, dsn)
		if err != nil {
			panic(err)
		} else if conn == nil {
			panic("empty database")
		}
	}

	return conn
}

func (x boxExec) Scan(rowsAffected *int, lastInsertID *int) (err error) {
	err = x.err
	if err != nil {
		return fmt.Errorf("database: BoxExec: %w", err)
	}

	if x.sqlResult == nil {
		return fmt.Errorf("database: BoxExec: %w", ErrInvalidArgumentsScan)
	}

	if rowsAffected != nil {
		if n, err := x.sqlResult.RowsAffected(); err == nil {
			if n < 1 {
				return fmt.Errorf("database: BoxExec: %w", sql.ErrNoRows)
			}

			*rowsAffected = int(n)
		}
	}

	if lastInsertID != nil {
		if n, err := x.sqlResult.LastInsertId(); err == nil {
			*lastInsertID = int(n)
		}
	}

	return err
}

func (SQL) SetupOrTeardown(ctx context.Context, conn ExecContext, queries ...string) error {
	for _, q := range queries {
		if _, err := conn.ExecContext(ctx, q); err != nil {
			return err
		}
	}

	return nil
}

// Scan the result of ExecContext that usually return numbers of rowsAffected
// and lastInsertID.
func (SQL) BoxExec(sqlResult sql.Result, err error) BoxExec { return boxExec{sqlResult, err} }

type BoxExec interface {
	Scan(rowsAffected *int, lastInsertID *int) (err error)
}

func (SQL) BoxQuery(sqlRows *sql.Rows, err error) BoxQuery { return boxQuery{sqlRows, err} }

type BoxQuery interface {
	Scan(row func(i int) entity.List) (err error)
}

type boxQuery struct {
	sqlRows *sql.Rows
	err     error
}

func (x boxQuery) Scan(row func(i int) entity.List) (err error) {
	err = x.err
	if err != nil {
		return err
	} else if x.sqlRows == nil {
		return fmt.Errorf("database: boxQuery: %w", sql.ErrNoRows)
	} else if err = x.sqlRows.Err(); err != nil {
		return err
	}
	defer x.sqlRows.Close()

	cols, err := x.sqlRows.Columns()
	if err != nil {
		return fmt.Errorf("database: boxQuery: %w", err)
	} else if len(cols) < 1 {
		return fmt.Errorf("database: boxQuery: %w", ErrNoColumnsReturned)
	}

	for i := 0; x.sqlRows.Next(); i++ {
		err = x.sqlRows.Err()
		if err != nil {
			return fmt.Errorf("database: boxQuery: %w", err)
		}

		dest := row(i)
		if dest == nil { // nil dest
			break
		} else if len(dest) < 1 { // empty dest
			continue
		} else if len(dest) != len(cols) { // diff dest & cols
			return fmt.Errorf("database: boxQuery: %w: [%d] columns on [%d] destinations", ErrInvalidArgumentsScan, len(cols), len(dest))
		}

		err = x.sqlRows.Scan(dest...) // scan into pointers
		if err != nil {
			return fmt.Errorf("database: boxQuery: %w", err)
		}
	}

	return err
}
