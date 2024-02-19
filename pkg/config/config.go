package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl              string
	AUTHTOKEN          string
	ACCOUNTSID         string
	SERVICESID         string
	ACCESS_KEY_ADMIN   string
	ACCESS_KEY_USER    string
	KEY_ID_FOR_PAY     string
	SECRET_KEY_FOR_PAY string
	PORT               string
}

func LoadEnvVariables() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, errors.New("failed to load .env variables")
	}

	config := Config{
		DBUrl:              os.Getenv("DB_URL"),
		AUTHTOKEN:          os.Getenv("DB_AUTHTOKEN"),
		ACCOUNTSID:         os.Getenv("DB_ACCOUNTSID"),
		SERVICESID:         os.Getenv("DB_SERVICESID"),
		ACCESS_KEY_ADMIN:   os.Getenv("ACCESS_KEY_ADMIN"),
		ACCESS_KEY_USER:    os.Getenv("ACCESS_KEY_USER"),
		KEY_ID_FOR_PAY:     os.Getenv("KEY_ID_FOR_PAY"),
		SECRET_KEY_FOR_PAY: os.Getenv("SECRET_KEY_FOR_PAY"),
		PORT:               os.Getenv("PORT"),
	}

	return config, nil
}
