package cmd

import (
	"time"

	. "github.com/Devil666face/gofinabot/config"
	"github.com/Devil666face/gofinabot/database"
	"github.com/Devil666face/gofinabot/models"
	. "github.com/Devil666face/gofinabot/routes"

	"github.com/vitaliy-ukiru/fsm-telebot"
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
	database.Storage()
	if dberr != nil {
		return nil, dberr
	}
	conf := telebot.Settings{
		Token:     Cfg.Token,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		Verbose:   Cfg.Debug,
		ParseMode: telebot.ModeHTML,
	}

	b, err := telebot.NewBot(conf)
	if err != nil {
		return nil, err
	}

	f := fsm.NewManager(b, nil, database.FsmStore, nil)
	SetMiddlewares(b, f)
	SetRoutes(b, f)

	return b, nil
}
