package handlers

import (
	"log"

	kb "github.com/Devil666face/gofinabot/markup"
	m "github.com/Devil666face/gofinabot/messages"

	. "github.com/Devil666face/gofinabot/models"

	telebot "gopkg.in/telebot.v3"
)

func delete(c telebot.Context) {
	if err := c.Delete(); err != nil {
		log.Print(err)
	}
}

func AskAdmins(c telebot.Context) {
	admins, err := GetAllAdmins()
	if err != nil {
		log.Print("Failed to get admins list")
	}
	for _, admin := range admins {
		if _, err := c.Bot().Send(&telebot.User{ID: int64(admin.TGID)}, m.AskAdminsForAddUser(c), kb.InlineAddUser(c.Chat().ID)); err != nil {
			log.Print(m.ErrSendMessage(admin.Username))
		}
	}
}

func Start(c telebot.Context) error {
	user := User{}
	if notfound := user.GetUserByTgID(c.Chat().ID); notfound != nil {
		user = User{
			TGID:     uint(c.Chat().ID),
			Username: c.Chat().Username,
		}
		if err := user.Create(); err != nil {
			return c.Send(m.ErrCreateUser(c))
		}
	}
	if !user.IsAllow {
		AskAdmins(c)
		return c.Send(m.SuccessfulCreateUser(c), telebot.RemoveKeyboard)
	}
	return c.Send(m.ErrUserAlreadyCreate(c), kb.Menu)
}
