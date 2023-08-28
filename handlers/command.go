package handlers

import (
	"log"

	. "github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils"
	"github.com/vitaliy-ukiru/fsm-telebot"

	telebot "gopkg.in/telebot.v3"
)

func delete(c telebot.Context) {
	err := c.Delete()
	if err != nil {
		log.Print(err)
	}
}

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
	}
	if !user.IsAllow {
		AskAdmins(c)
		return c.Send(utils.SuccessfulCreateUser(c), telebot.RemoveKeyboard)
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
	err = c.Delete()
	if err != nil {
		log.Print(err)
	}
	_, err = c.Bot().Send(&telebot.User{ID: int64(user.TGID)}, utils.PermissionsForUserAdded(user.Username), utils.Menu)
	if err != nil {
		return c.Send(utils.ErrSendMessage(user.Username))
	}
	return c.Send(utils.SuccessfulUpdateUser(user.Username), utils.Menu)
}

func OnIgnoreUser(c telebot.Context) error {
	return c.Delete()
}

func OnBackBtn(c telebot.Context, f fsm.Context) error {
	go f.Finish(true)
	return c.Send(utils.GO_BACK, utils.Menu)
}
