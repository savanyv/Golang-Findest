package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		PGHost     string
		PGPort     string
		PGUser     string
		PGPassword string
		PGDatabase string
	}
	Jwt struct {
		SecretKey string
	}
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

	return &Config{
		Database: struct{PGHost, PGPort, PGUser, PGPassword, PGDatabase string}{
			PGHost:     getEnv("PGHOST"),
			PGPort:     getEnv("PGPORT"),
			PGUser:     getEnv("PGUSER"),
			PGPassword: getEnv("PGPASSWORD"),
			PGDatabase: getEnv("PGDATABASE"),
		},
		Jwt: struct{SecretKey string}{
			SecretKey: getEnv("SECRETKEY"),
		},
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Panic(key + " is not set")
	}
	return value
}