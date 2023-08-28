package handlers

import (
	. "github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

var (
	InputType     = fsm.NewStateGroup("type")
	InputTypeName = InputType.New("type:name")
)

func OnTransBtn(c telebot.Context) error {
	return c.Send(utils.ChangeTypeForUpdate(), utils.TypeMenu)
}

func OnTypeAddBtn(c telebot.Context, s fsm.Context) error {
	s.Set(InputTypeName)
	return c.Send(utils.AddNewType())
}

func OnTypeNameRecive(c telebot.Context, s fsm.Context) error {
	go s.Finish(true)
	trtype := TypeTransaction{
		Type:   c.Message().Text,
		UserID: uint(c.Chat().ID),
	}
	err := trtype.Create()
	if err != nil {
		return c.Send(utils.ErrCreateType(c))
	}
	return c.Send(utils.SuccessfulCreateType(c), utils.Menu)
}
