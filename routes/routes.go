package routes

import (
	. "github.com/Devil666face/gofinabot/handlers"
	"github.com/Devil666face/gofinabot/utils"

	telebot "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func SetMiddlewares(b *telebot.Bot) {
	b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())
}

func SetRoutes(b *telebot.Bot) {
	b.Handle("/start", Start)
	b.Handle(&utils.NewTransBtn, TransCreate)
}
