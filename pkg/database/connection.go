package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(connURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
