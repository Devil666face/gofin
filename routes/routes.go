package routes

import (
	. "github.com/Devil666face/gofinabot/config"
	. "github.com/Devil666face/gofinabot/handlers"
	kb "github.com/Devil666face/gofinabot/markup"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func SetMiddlewares(b *telebot.Bot, f *fsm.Manager) {
	if Cfg.Log {
		b.Use(middleware.Logger())
	}
	b.Use(middleware.AutoRespond())
}

func SetRoutes(b *telebot.Bot, f *fsm.Manager) {

	allowGroup := b.Group()
	allowGroupFsm := f
	allowGroup.Use(AllowOnlyMiddleware)
	allowGroupFsm.Use(AllowOnlyMiddleware)

	setFreeRoutes(b, f)
	setCallbackRoutes(allowGroup, allowGroupFsm)
	setTypeRoutes(allowGroup, allowGroupFsm)
	setTransRoutes(allowGroup, allowGroupFsm)
	setStatisticRoutes(allowGroup, allowGroupFsm)
}

func setFreeRoutes(b *telebot.Bot, f *fsm.Manager) {
	b.Handle("/start", Start)
}

func setCallbackRoutes(b *telebot.Group, f *fsm.Manager) {
	f.Bind(telebot.OnCallback, fsm.AnyState, CallbackHandler, CallbackKeyValueMiddleware)
	f.Bind(&kb.BackBtn, fsm.AnyState, OnBackBtn)
}

func setTypeRoutes(b *telebot.Group, f *fsm.Manager) {
	b.Handle(&kb.TypeTransBtn, OnTypeBtn)
	f.Bind(&kb.TypeAddBtn, fsm.DefaultState, OnTypeAddBtn)
	f.Bind(telebot.OnText, InputTypeNameState, OnTypeNameRecive)
	f.Bind(telebot.OnText, InputTypeNameForUpdateState, OnTypeNameForUpdateRecive)
}

func setTransRoutes(b *telebot.Group, f *fsm.Manager) {
	f.Bind(&kb.AddTransBtn, fsm.DefaultState, OnAddTransBtn)
	f.Bind(telebot.OnText, InputTransCommentState, OnTransCommentRecive)
	f.Bind(telebot.OnText, InputTransValueState, OnTransValueRecive)
}

func setStatisticRoutes(b *telebot.Group, f *fsm.Manager) {
	b.Handle(&kb.MonthStatBtn, OnMonthStatBtn)
	b.Handle(&kb.ExcelReportBtn, OnExcelReportBtn)
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
	case kb.CALLBACK_TRANS_BALANCE:
		return OnTransBalanceInlineBtn(c, s)
	case kb.CALLBACK_TRANS_TYPE:
		return OnTransTypeInlibeBtn(c, s)
	case kb.CALLBACK_TRANS_EMPTY_COMMENT:
		return OnTransCommentRecive(c, s)
	case kb.CALLBACK_TRANS_CREATE:
		return OnTransAddInlineBtn(c, s)
	case kb.CALLBACK_TRANS_STAT:
		return OnTransStatInlineBtn(c)
	case kb.CALLBACK_TRANS_DELETE:
		return OnTransDeleteInlineBtn(c)
	}
	return nil
}
