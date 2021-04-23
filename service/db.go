package service

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbHandle *sql.DB

func GetDB() (*sql.DB, error) {
	if dbHandle != nil {
		return dbHandle, nil
	}
	dsn := os.Getenv("DB_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	dbHandle = db
	return db, nil
}
