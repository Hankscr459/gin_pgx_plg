package configs

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
)

func ConnectDb() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("PostgresURI"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
	// defer conn.Close(context.Background())
}

func Migrate() error {

	db, err := sql.Open("postgres", os.Getenv("PostgresURI"))
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://plugins/configs/database/migrations",
		"postgres", driver)
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		return err
	}
	return nil
}
