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
	// b.Handle(telebot.OnCallback, CallbackHandler, CallbackKeyValueMw)

	allow := b.Group()
	fsmallow := f.NewGroup()

	fsmallow.Use(AllowOnlyMw)
	fsmallow.Bind(telebot.OnCallback, fsm.AnyState, CallbackHandler, CallbackKeyValueMw)

	allow.Use(AllowOnlyMw)

	allow.Handle(&utils.TypeTransBtn, OnTransBtn)
	fsmallow.Bind(&utils.BackButton, fsm.AnyState, OnBackBtn)
	fsmallow.Bind(&utils.TypeAddButton, fsm.DefaultState, OnTypeAddBtn)

	fsmallow.Bind(telebot.OnText, InputTypeName, OnTypeNameRecive)
	fsmallow.Bind(telebot.OnText, InputTypeNameForUpdate, OnTypeNameForUpdateRecive)
}

func CallbackHandler(c telebot.Context, s fsm.Context) error {
	switch c.Get(CALLBACK_KEY) {
	case utils.CALLBACK_CONFIRM_USER:
		return AllowOnlyMw(OnConfirmUser)(c)
	case utils.CALLBACK_IGNORE_USER:
		return AllowOnlyMw(OnIgnoreUser)(c)
	case utils.CALLBACK_TYPE:
		return AllowOnlyMw(OnUpdateTypeInlineBtn)(c)
	case utils.CALLBACK_TYPE_UPDATE:
		return OnUpdateTypeNameBtn(c, s)
	case utils.CALLBACK_TYPE_DELETE:
		return AllowOnlyMw(OnDeleteTypeBtn)(c)
	}
	return nil
}
