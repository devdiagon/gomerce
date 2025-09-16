package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost       string
	Port             string
	DBUser           string
	DBPassword       string
	DBADDRESS        string
	DBNAME           string
	JWTExpirationSec int64
	JWTSecret        string
}

var Envs = initConfig()

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

func getEnvHasInt(key string, fallback int64) int64 {
	value, ok := os.LookupEnv(key)

	if ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:       getEnv("PUBLIC_HOST", "http://localhost"),
		Port:             fmt.Sprintf(":%s", getEnv("PORT", "8080")),
		DBUser:           getEnv("DB_USER", "admin"),
		DBPassword:       getEnv("DB_PASSWORD", "admin"),
		DBADDRESS:        fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBNAME:           getEnv("DB_NAME", "gomerce"),
		JWTExpirationSec: getEnvHasInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecret:        getEnv("JWT_SECRET", "s3CreT157ZXY"),
	}
}
