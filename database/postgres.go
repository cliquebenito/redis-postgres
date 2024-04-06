package database

import (
	"CRUD-REDIS/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PostgresConnect(config *config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBhost, config.DBport, config.DBuser, config.DBpass, config.DBname)
	db, _ := sql.Open("postgres", psqlInfo)

	fmt.Println("Successfully connected to Postgres!")
	return db
}
