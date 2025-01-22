package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConnection struct {
	DbConn *pgxpool.Pool
}

var Db DatabaseConnection

func Connect() error {
	conn, err := pgxpool.New(context.Background(), os.Getenv("CONN_URL"))
	if err != nil {
		return err
	}

	Db.DbConn = conn
	return nil
}
