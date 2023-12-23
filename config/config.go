package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

var conf cfg
var dbConn *sql.DB

type cfg struct {
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	WebPort    string `mapstructure:"WEB_PORT"`
}

func LoadConfig() (*cfg, error) {	
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func ConnectToDb(conf *cfg) error {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName)
	db, _ := sql.Open("postgres", connStr)
	err := db.Ping()

	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS pessoa (id UUID PRIMARY KEY,apelido VARCHAR(32) UNIQUE NOT NULL,nome VARCHAR(100) NOT NULL,nascimento DATE NOT NULL, stack JSON)")

	return err
}

func GetDBConn() *sql.DB {
	return dbConn
}
