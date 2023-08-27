package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	TOKEN = env("TOKEN", "")
	DB    = env("DB", "db.sqlite3")
	DEBUG = boolenv(env("DEBUG", "False"))
)

func env(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	if fallback != "" {
		return fallback
	}
	panic(fmt.Sprintf(`Environment variable not found %v`, name))
}

func boolenv(env string) bool {
	if env == "True" || env == "true" {
		return true
	}
	return false
}
