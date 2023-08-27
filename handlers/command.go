package handlers

import (
	"fmt"
	"log"

	. "github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils"

	telebot "gopkg.in/telebot.v3"
)

func AskAdmins(c telebot.Context) {
	admins, err := GetAllAdmins()
	if err != nil {
		log.Print("Failed to get admins list")
	}
	for _, admin := range admins {
		_, err := c.Bot().Send(&telebot.User{ID: int64(admin.TGID)}, utils.AskAdminsForAddUser(c), utils.InlineAddUser(c.Chat().ID))
		if err != nil {
			log.Print(utils.ErrSendMessage(admin.Username))
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
		err := user.Create()
		if err != nil {
			return c.Send(utils.ErrCreateUser(c))
		}
		AskAdmins(c)
		return c.Send(utils.SuccessfulCreateUser(c), utils.Menu)
	}
	if !user.IsAllow {
		AskAdmins(c)
	}
	return c.Send(utils.UserAlreadyCreate(c), utils.Menu)
}

func OnConfirmUser(c telebot.Context) error {
	id := utils.ToInt64(c.Get(CALLBACK_VAL))
	user := User{}
	if notfound := user.GetUserByTgID(id); notfound != nil {
		return c.Send(utils.UserNotFound(id))
	}
	user.IsAllow = true
	err := user.Update()
	if err != nil {
		return c.Send(utils.ErrUserUpdate(user.Username))
	}
	return c.Send(utils.SuccessfulUpdateUser(user.Username))
}

func OnIgnoreUser(c telebot.Context) error {
	fmt.Println("Ignore")
	return nil
}

func TransCreate(c telebot.Context) error {
	return c.Send("FOO", telebot.RemoveKeyboard)
}
