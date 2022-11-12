package main

import (
	"fmt"
	"log"

	api "github.com/book/api"
	config "github.com/book/config"
	"github.com/book/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "net/http"
)

func main() {
	connSqlx, err := sqlx.Connect("postgres", config.ConnStr())
	if err != nil {
		fmt.Printf("sqlx ulanishda xatolik(main func'da \n%v)", err)
		return
	}

	db := postgres.NewBookRepo(connSqlx)

	server := api.NewServer(db)

	err = server.Run()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
