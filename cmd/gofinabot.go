package cmd

import (
	"time"

	"github.com/Devil666face/gofinabot/config"
	"github.com/Devil666face/gofinabot/database"
	"github.com/Devil666face/gofinabot/models"

	telebot "gopkg.in/telebot.v3"
)

func Migrate() error {
	err := database.Migrate(
		&models.MoneyTransaction{},
		&models.TypeTransaction{},
		&models.User{},
	)
	if err != nil {
		return err
	}
	return nil
}

func Bot() (*telebot.Bot, error) {
	dberr := database.Connect()
	if dberr != nil {
		return nil, dberr
	}
	conf := telebot.Settings{
		Token:  config.TOKEN,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(conf)
	if err != nil {
		return nil, err
	}

	return b, nil
}
