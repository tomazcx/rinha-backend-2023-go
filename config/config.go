package config

import (
	"github.com/spf13/viper"
)

var conf Cfg

type Cfg struct {
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBHost         string `mapstructure:"DB_HOST"`
	WebPort        string `mapstructure:"WEB_PORT"`
}

func LoadConfig() (*Cfg, error) {
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
