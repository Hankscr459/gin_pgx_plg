package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
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
