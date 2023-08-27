package handlers

import (
	. "github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils"

	telebot "gopkg.in/telebot.v3"
)

func Start(c telebot.Context) error {
	user := User{}
	if notfound := user.GetUserByTgID(c.Chat().ID); notfound != nil {
		user = User{
			TGID:     uint(c.Chat().ID),
			Username: c.Chat().Username,
			IsAdmin:  false,
		}
		err := user.Create()
		if err != nil {
			return c.Send(utils.ErrCreateUser(c))
		}
		return c.Send(utils.SuccessfulCreateUser(c), utils.Menu)
	}
	return c.Send(utils.UserAlreadyCreate(c), utils.Menu)
}

func TransCreate(c telebot.Context) error {
	return c.Send("FOO", utils.RemoveMenu)
}
