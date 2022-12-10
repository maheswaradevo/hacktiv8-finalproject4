package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/config"
)

func GetDatabase() *sql.DB {
	log.Printf("INFO GetDatabase database connection: starting database connection process")

	cfg := config.GetConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Address, cfg.Database.Name)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error GetDatabase sql open connection fatal error: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("ERROR GetDatabase db ping fatal error: %v", err)
	}
	log.Printf("INFO GetDatabase database connectionn: established successfully\n")
	return db
}
