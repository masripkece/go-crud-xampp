package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/produkdb")
	if err != nil {
		log.Fatal("Gagal koneksi DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB tidak merespons:", err)
	}
}
