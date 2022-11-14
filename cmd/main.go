package main

import (
	"fmt"
	"log"
	_ "net/http"

	"github.com/book/api"
	"github.com/book/config"
	"github.com/book/storage/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load("./..")
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	storage := postgres.NewBookRepo(db)
	server := api.NewServer(storage)
	fmt.Println(cfg.HttpPort, "--------------------------------------------")
	err = server.Run(":8000")
	
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
