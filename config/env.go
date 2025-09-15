package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBADDRESS  string
	DBNAME     string
}

var Envs = initConfig()

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       fmt.Sprintf(":%s", getEnv("PORT", "8080")),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", "admin"),
		DBADDRESS:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBNAME:     getEnv("DB_NAME", "gomerce"),
	}
}
