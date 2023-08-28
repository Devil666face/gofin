package routes

import (
	. "github.com/Devil666face/gofinabot/handlers"
	"github.com/Devil666face/gofinabot/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func SetMiddlewares(b *telebot.Bot, f *fsm.Manager) {
	// b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())
}

func SetRoutes(b *telebot.Bot, f *fsm.Manager) {
	b.Handle("/start", Start)
	b.Handle(telebot.OnCallback, CallbackHandler, CallbackKeyValueMw)

	allow := b.Group()
	fsmallow := f.NewGroup()

	allow.Use(AllowOnlyMw)
	fsmallow.Use(AllowOnlyMw)

	allow.Handle(&utils.TypeTransBtn, OnTransBtn)
	fsmallow.Bind(&utils.BackButton, fsm.AnyState, OnBackBtn)

	fsmallow.Bind(&utils.TypeAddButton, fsm.DefaultState, OnTypeAddBtn)
	fsmallow.Bind(telebot.OnText, InputTypeName, OnTypeNameRecive)
}

func CallbackHandler(c telebot.Context) error {
	switch c.Get(CALLBACK_KEY) {
	case utils.CALLBACK_CONFIRM_USER:
		return AllowOnlyMw(OnConfirmUser)(c)
	case utils.CALLBACK_IGNORE_USER:
		return AllowOnlyMw(OnIgnoreUser)(c)
	}
	return nil
}
