package cmd

import (
	"time"

	. "github.com/Devil666face/gofinabot/config"
	"github.com/Devil666face/gofinabot/models"
	. "github.com/Devil666face/gofinabot/routes"
	"github.com/Devil666face/gofinabot/store/database"
	"github.com/Devil666face/gofinabot/store/memstore"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

func Migrate() error {
	if err := database.Migrate(
		&models.MoneyTransaction{},
		&models.TypeTransaction{},
		&models.User{},
	); err != nil {
		return err
	}
	return nil
}

func Bot() (*telebot.Bot, error) {
	memstore.Store()
	if err := database.Connect(); err != nil {
		return nil, err
	}
	if err := Migrate(); err != nil {
		return nil, err
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

	f := fsm.NewManager(b, nil, memstore.Memstore, nil)

	SetMiddlewares(b, f)
	SetRoutes(b, f)

	return b, nil
}
