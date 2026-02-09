package psql

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

type PsqlDB struct {
	db *pgx.Pool
}

func (p *PsqlDB) Conn() *pgx.Pool {
	return p.db
}

func New() *PsqlDB {
		urlExample := "postgres://myuser:secret@localhost:5431/ecommerce_db"
	db, err := pgx.New(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("OK")

	return &PsqlDB{db: db}

	
}