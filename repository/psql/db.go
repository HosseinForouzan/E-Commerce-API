package psql

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Username string	`koanf:"username"`
	Password string	`koanf:"password"`
	Port int	`koanf:"port"`
	Host string	`koanf:"host"`
	DBName string	`koanf:"db_name"`
} 

type PsqlDB struct {
	config Config
	db *pgx.Pool
}

func (p *PsqlDB) Conn() *pgx.Pool {
	return p.db
}

func New(config Config) *PsqlDB {
	// urlExample := "postgres://myuser:secret@localhost:5431/ecommerce_db"
	urlExample := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)
	db, err := pgx.New(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("OK")

	return &PsqlDB{db: db, config: config}

	
}