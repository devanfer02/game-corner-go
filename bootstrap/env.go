package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv               string `mapstructure:"APP_ENV"`
	ServerAddress        string `mapstructure:"SERVER_ADDRESS"`
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASSWORD"`
	DBName               string `mapstructure:"DB_NAME"`
	AccessTokenExpyHour  string `mapstructure:"ACCESS_TOKEN_EXPY_HOUR"`
	RefreshTokenExpyHour string `mapstructure:"REFRESH_TOKEN_EXPY_HOUR"`
	AccessTokenSecret    string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret   string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

var globenv *Env

func InitEnv() {
	env := Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Can't find the env file. ERR: %s\n", err.Error())
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("Environment variables can't be loaded, ERR: %s\n", err.Error())
	}

	if env.AppEnv == "devlopment" {
		log.Println("Server App is running on development env")
	}

	globenv = &env 
}