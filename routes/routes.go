package routes

import (
	. "github.com/Devil666face/gofinabot/handlers"
	kb "github.com/Devil666face/gofinabot/markup"

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

	fsmallow.Use(AllowOnlyMiddleware)
	allow.Use(AllowOnlyMiddleware)

	fsmallow.Bind(telebot.OnCallback, fsm.AnyState, CallbackHandler, CallbackKeyValueMiddleware)

	allow.Handle(&kb.TypeTransBtn, OnTransBtn)
	fsmallow.Bind(&kb.BackButton, fsm.AnyState, OnBackBtn)
	fsmallow.Bind(&kb.TypeAddButton, fsm.DefaultState, OnTypeAddBtn)

	fsmallow.Bind(telebot.OnText, InputTypeNameState, OnTypeNameRecive)
	fsmallow.Bind(telebot.OnText, InputTypeNameForUpdateState, OnTypeNameForUpdateRecive)
}

func CallbackHandler(c telebot.Context, s fsm.Context) error {
	switch c.Get(CALLBACK_KEY) {
	case kb.CALLBACK_CONFIRM_USER:
		return AdminOnlyMiddleware(OnConfirmUser)(c)
	case kb.CALLBACK_IGNORE_USER:
		return AdminOnlyMiddleware(OnIgnoreUser)(c)
	case kb.CALLBACK_TYPE:
		return OnUpdateCurrentTypeInlineBtn(c)
	case kb.CALLBACK_TYPE_UPDATE:
		return OnUpdateTypeNameInlineBtn(c, s)
	case kb.CALLBACK_TYPE_DELETE:
		return OnDeleteTypeInlineBtn(c)
	}
	return nil
}
