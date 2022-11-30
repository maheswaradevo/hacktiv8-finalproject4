package config

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	Database      Database
	ServerAddress string
	WhiteListed   string

	JWT_SIGNING_METHOD jwt.SigningMethod
	API_SECRET_KEY     string
}
type Database struct {
	Username string
	Password string
	Address  string
	Port     string
	Name     string
}

var config Config

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("ERROR .env Not found")
	}

	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	config.Port = os.Getenv("PORT")
	config.Database.Username = os.Getenv("DB_USERNAME")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.Address = os.Getenv("DB_ADDRESS")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.Name = os.Getenv("DB_NAME")
	config.WhiteListed = os.Getenv("WHITELISTED_URLS")

	config.JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	config.API_SECRET_KEY = os.Getenv("SECRET_KEY")
}

func GetConfig() *Config {
	return &config
}
