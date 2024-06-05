package config

import (
    "context"
    "github.com/jackc/pgx/v4/pgxpool"
    "log"
)

var DB *pgxpool.Pool

func InitDB() {
    var err error
    connString := "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"
    DB, err = pgxpool.Connect(context.Background(), connString)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    } else {
        log.Println("Database connection established")
    }
}
