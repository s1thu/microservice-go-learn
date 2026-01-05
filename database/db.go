package database

import (
	"database/sql"
	"example/go-web-gin/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	POSTGRES = "postgres"
	MYSQL    = "mysql"
)

func ConnectDB() {
	cfg := config.AppConfig

	var db *sql.DB
	var err error

	switch cfg.DBDRIVER {
	case POSTGRES:
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
		)
		db, err = sql.Open(POSTGRES, dsn)

	case MYSQL:
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
		)
		db, err = sql.Open(MYSQL, dsn)
	default:
		log.Fatal("Unsupported DB driver:", cfg.DBDRIVER)
	}
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connection established")

}
