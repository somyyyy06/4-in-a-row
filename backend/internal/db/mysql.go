package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQL(dsn string) {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open MySQL:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to connect MySQL:", err)
	}

	log.Println("MySQL connected")
}
