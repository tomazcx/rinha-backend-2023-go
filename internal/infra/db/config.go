package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/tomazcx/rinha-backend-go/config"
)

var dbConn *sql.DB

func ConnectToDb(conf *config.Cfg) error {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName)
	log.Println(connStr)
	db, _ := sql.Open("postgres", connStr)
	err := db.Ping()

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(150)
	db.SetMaxIdleConns(150)

	dbConn = db

	return nil
}

func GetDBConn() *sql.DB {
	return dbConn
}
