package migrator

import (
	"database/sql"
	"fmt"

	"github.com/rubenv/sql-migrate"
)

type Migrator struct {
	migrations *migrate.FileMigrationSource
}

func New() Migrator {
	migrations := &migrate.FileMigrationSource{
    Dir: "./repository/psql/migrations",
}
	return Migrator{migrations: migrations}
}

func (m Migrator) Up() {
		filename := "postgres://myuser:secret@localhost:5431/ecommerce_db"
	db, err := sql.Open("postgres", filename)
	if err != nil {
		panic(fmt.Errorf("can't open database. %w", err))
	}

	n, err := migrate.Exec(db, "postgres", m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't execute migration %w", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	filename := "postgres://myuser:secret@localhost:5431/ecommerce_db"
	db, err := sql.Open("postgres", filename)
	if err != nil {
		panic(fmt.Errorf("can't open database %w", err))
	}

	n, err := migrate.Exec(db, "postgres", m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't rollback migration %w", err))
	}
	fmt.Printf("Rollback %d migrations!\n", n)
}

func (m Migrator) Status() {
	filename := "postgres://myuser:secret@localhost:5431/ecommerce_db"
	db, err := sql.Open("postgres", filename)
	if err != nil {
		// Handle errors!
	}

	n, err := migrate.Exec(db, "postgres", m.migrations, migrate.Up)
	if err != nil {
		// Handle errors!
	}
	fmt.Printf("Applied %d migrations!\n", n)

}