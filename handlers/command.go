package handlers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
)

func Start(c telebot.Context) error {
	tgid := c.Chat().ID
	return c.Send(fmt.Sprintf("%d", tgid))
}
