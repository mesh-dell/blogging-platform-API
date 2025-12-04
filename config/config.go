package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBAddr     string
	Port       string
}

func GetConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	c := Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBAddr:     os.Getenv("DB_ADDR"),
		Port:       os.Getenv("PORT"),
	}

	if c.DBName == "" || c.DBUser == "" || c.DBPassword == "" {
		return c, errors.New("missing DB_USER, DB_PASSWORD or DB_NAME environment variables")
	}

	if c.Port == "" {
		c.Port = "8080"
	}

	if c.DBAddr == "" {
		c.DBAddr = "localhost:3306"
	}
	return c, nil
}
