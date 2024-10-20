// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func New(db DBTX) *Connection {
	return &Connection{
		ctx: context.Background(),
		db:  db,
	}
}

type Connection struct {
	ctx context.Context
	db  DBTX
}

func (q *Connection) WithTx(tx pgx.Tx) *Connection {
	return &Connection{
		db: tx,
	}
}
