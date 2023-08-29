package handlers

import (
	"log"
	"strings"

	. "github.com/Devil666face/gofinabot/models"

	telebot "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

var (
	CALLBACK_KEY = "callback_key"
	CALLBACK_VAL = "callback_val"
)

func CallbackKeyValueMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Callback() != nil {
			r := strings.ReplaceAll(c.Callback().Data, "\f", "")
			data := strings.Split(r, ":")
			c.Set(CALLBACK_KEY, data[0])
			c.Set(CALLBACK_VAL, data[1])
		}
		return next(c)
	}
}

func permission(selectorFunc func() ([]User, error), next telebot.HandlerFunc) telebot.HandlerFunc {
	chats, err := GetChatIdsForSelector(selectorFunc)
	if err != nil {
		log.Print(err)
		return middleware.Whitelist()(next)
	}
	return middleware.Whitelist(chats...)(next)

}

func AdminOnlyMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return permission(GetAllAdmins, next)
}

func AllowOnlyMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return permission(GetAllAllows, next)
}
