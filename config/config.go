package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Token       string `env:"TOKEN" env-required:"true"`
	Db          string `env:"DB" env-default:"db.sqlite3"`
	Debug       bool   `env:"DEBUG" env-default:"false"`
	SuperuserId uint   `env:"SUPERUSERID" env-default:"446545799"`
	Log         bool   `env:"LOG" env-default:"true"`
}

var Cfg Config

func init() {
	err := cleanenv.ReadEnv(&Cfg)
	if err != nil {
		panic(err)
	}
}
