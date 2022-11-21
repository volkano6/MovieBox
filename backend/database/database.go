package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", os.Getenv("DB"))
	if err != nil {
		panic(err.Error())
	}
	DB = db
	fmt.Println("Successfully connected to the database.")
}
