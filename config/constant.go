package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	APP_PORT,
	DB_HOST,
	DB_USER,
	DB_PASSWORD,
	DB_NAME,
	DB_PORT,
	DB_SSLMODE,
	DB_TIMEZONE,
	SECRET_KEY string
)

func InitializeConstantValue() {
	APP_PORT = os.Getenv("APP_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_SSLMODE = os.Getenv("DB_SSLMODE")
	DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
	SECRET_KEY = os.Getenv("SECRET_KEY")
}
