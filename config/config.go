package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Token string `env:"TOKEN" env-required:"true"`
	Db    string `env:"DB" env-default:"db.sqlite3"`
	Debug bool   `env:"DEBUG" env-default:"false"`
}

var Cfg Config

func init() {
	err := cleanenv.ReadEnv(&Cfg)
	if err != nil {
		panic(err)
	}
}
