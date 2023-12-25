package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tomazcx/rinha-backend-go/config"
)

var dbConn *sql.DB

func ConnectToDb(conf *config.Cfg) error {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName)
	db, _ := sql.Open("postgres", connStr)
	err := db.Ping()

	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS pessoa (id UUID PRIMARY KEY,apelido VARCHAR(32) UNIQUE NOT NULL,nome VARCHAR(100) NOT NULL, nascimento DATE NOT NULL, stack TEXT)")

	return err
}

func GetDBConn() *sql.DB {
	return dbConn
}
