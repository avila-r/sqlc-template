package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"

	"github.com/avila-r/env"
	"github.com/avila-r/tasker"
)

var (
	Conn = func() *Connection {
		if err := env.Load(tasker.RootPath); err != nil {
			log.Fatalf("error while loading .env file - %v", err.Error())
		}

		ctx, dsn := context.Background(), env.Get("POSTGRES_DSN")

		conn, err := pgx.Connect(ctx, dsn)

		if err != nil {
			log.Fatalf("error while estabilishing conn with postgres - %v", err.Error())
		}

		return New(conn)
	}()
)
