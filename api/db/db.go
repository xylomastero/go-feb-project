package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Initialize() error {
	var err error
	DB, err = sql.Open("mysql", "xylomastero:student@tcp(127.0.0.1:3306)/car_db")
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Connected to database")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
