package handlers

import (
	// "log"

	// . "github.com/Devil666face/gofinabot/models"
	// "github.com/Devil666face/gofinabot/utils"

	"strings"

	telebot "gopkg.in/telebot.v3"
)

var (
	CALLBACK_KEY = "callback_key"
	CALLBACK_VAL = "callback_val"
)

func CallbackKeyValueMw(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Callback() != nil {
			r := strings.ReplaceAll(c.Callback().Data, "\f", "")
			data := strings.Split(r, ":")
			c.Set(CALLBACK_KEY, data[0])
			c.Set(CALLBACK_VAL, data[1])
		}
		return next(c) // continue execution chain
	}
}
