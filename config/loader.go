package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		JWT: JWT{
			Secret: os.Getenv("JWT_SECRET"),
			Expire: os.Getenv("JWT_EXPIRE"),
		},
	}
}

func LoadTestConfig() *Config {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("APP_TEST_HOST"),
			Port: os.Getenv("APP_TEST_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_TEST_HOST"),
			Port:     os.Getenv("DB_TEST_PORT"),
			Username: os.Getenv("DB_TEST_USERNAME"),
			Password: os.Getenv("DB_TEST_PASSWORD"),
			Name:     os.Getenv("DB_TEST_NAME"),
		},
		JWT: JWT{
			Secret: os.Getenv("JWT_TEST_SECRET"),
			Expire: os.Getenv("JWT_TEST_EXPIRE"),
		},
	}
}
