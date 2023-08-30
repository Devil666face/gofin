package markup

import (
	"fmt"

	. "github.com/Devil666face/gofinabot/messages"
	"github.com/Devil666face/gofinabot/models"

	telebot "gopkg.in/telebot.v3"
)

var (
	AddTransBtn = telebot.ReplyButton{
		Text: TRANS_NEW,
	}
	TypeTransBtn = telebot.ReplyButton{
		Text: TRANS_TYPE,
	}
	MonthStatBtn = telebot.ReplyButton{
		Text: STATISTIC_TRANS_MONTH,
	}
	Menu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{AddTransBtn},
			{TypeTransBtn},
			{MonthStatBtn},
		},
		ResizeKeyboard: true,
	}
)

var (
	TypeMenu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{TypeAddBtn},
			{BackBtn},
		},
		ResizeKeyboard: true,
	}
	TypeAddBtn = telebot.ReplyButton{
		Text: TYPE_ADD,
	}
	BackBtn = telebot.ReplyButton{
		Text: BACK,
	}
)

var (
	TransMenu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{BackBtn},
		},
		ResizeKeyboard: true,
	}
	BalanceInline = &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{ExpenseInlineBtn, IncomeInlineBtn},
		},
		ResizeKeyboard: true,
	}
	ExpenseInlineBtn = telebot.InlineButton{
		Text:   EXPENSE,
		Unique: fmt.Sprintf("%s:%t", CALLBACK_TRANS_BALANCE, CALLBACK_EXPENSE),
	}
	IncomeInlineBtn = telebot.InlineButton{
		Text:   INCOME,
		Unique: fmt.Sprintf("%s:%t", CALLBACK_TRANS_BALANCE, CALLBACK_INCOME),
	}
	TransCommentInline = &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{telebot.InlineButton{
				Text:   EMPTY_COMMENT,
				Unique: fmt.Sprintf("%s:%s", CALLBACK_TRANS_EMPTY_COMMENT, "empty"),
			}},
		},
		ResizeKeyboard: true,
	}
	CreateTransInlineBtn = telebot.InlineButton{
		Text:   TRANS_CREATE,
		Unique: fmt.Sprintf("%s:%t", CALLBACK_TRANS_CREATE, CALLBACK_TRANS_CREATE_ADD),
	}
	CancelTransInlineBtn = telebot.InlineButton{
		Text:   TRANS_CANCEL,
		Unique: fmt.Sprintf("%s:%t", CALLBACK_TRANS_CREATE, CALLBACK_TRANS_CREATE_CANCEL),
	}
	TransAddInline = &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{CreateTransInlineBtn, CancelTransInlineBtn},
		},
		ResizeKeyboard: true,
	}
)

var (
	StatMenu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{ExcelReportBtn},
			{BackBtn},
		},
		ResizeKeyboard: true,
	}
	ExcelReportBtn = telebot.ReplyButton{
		Text: STATISTIC_EXCEL_REPORT,
	}
)

var (
	CALLBACK_CONFIRM_USER = "confirm_user"
	CALLBACK_IGNORE_USER  = "ignore_user"
)

var (
	CALLBACK_TYPE        = "type"
	CALLBACK_TYPE_UPDATE = "type_update"
	CALLBACK_TYPE_DELETE = "type_delete"
)

var (
	CALLBACK_TRANS_BALANCE       = "trans_balance"
	CALLBACK_EXPENSE             = false
	CALLBACK_INCOME              = true
	CALLBACK_TRANS_TYPE          = "trans_type"
	CALLBACK_TRANS_EMPTY_COMMENT = "trans_comment"
	CALLBACK_TRANS_CREATE        = "trans_create"
	CALLBACK_TRANS_CREATE_ADD    = true
	CALLBACK_TRANS_CREATE_CANCEL = false
)

var (
	CALLBACK_TRANS_STAT   = "trans"
	CALLBACK_TRANS_DELETE = "trans_delete"
)

func InlineAddUser(id int64) *telebot.ReplyMarkup {
	confirm := telebot.InlineButton{Text: CONFIRM_USER, Unique: fmt.Sprintf("%s:%d", CALLBACK_CONFIRM_USER, id)}
	ignore := telebot.InlineButton{Text: IGNORE_USER, Unique: fmt.Sprintf("%s:%d", CALLBACK_IGNORE_USER, id)}
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{confirm, ignore},
		},
		ResizeKeyboard: true,
	}
}

func inlineTypesWithCallback(trtypes []models.TypeTransaction, callbackey string) *telebot.ReplyMarkup {
	btns := [][]telebot.InlineButton{}
	for _, v := range trtypes {
		btn := telebot.InlineButton{Text: v.Type, Unique: fmt.Sprintf("%s:%d", callbackey, v.ID)}
		btns = append(btns, []telebot.InlineButton{btn})
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: btns,
		ResizeKeyboard: true,
	}
}

func InlineTypes(trtypes []models.TypeTransaction) *telebot.ReplyMarkup {
	return inlineTypesWithCallback(trtypes, CALLBACK_TYPE)
}

func InlineTypesForAddTrans(trtypes []models.TypeTransaction) *telebot.ReplyMarkup {
	return inlineTypesWithCallback(trtypes, CALLBACK_TRANS_TYPE)
}

func UpdateTypeInline(trtype models.TypeTransaction) *telebot.ReplyMarkup {
	update := telebot.InlineButton{Text: UpdateTypeText(trtype.Type), Unique: fmt.Sprintf("%s:%d", CALLBACK_TYPE_UPDATE, trtype.ID)}
	delete := telebot.InlineButton{Text: DeleteTypeText(trtype.Type), Unique: fmt.Sprintf("%s:%d", CALLBACK_TYPE_DELETE, trtype.ID)}
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{update},
			{delete},
		},
		ResizeKeyboard: true,
	}
}

func InlineTransStatList(transes []models.MoneyTransaction) *telebot.ReplyMarkup {
	inlinekb := [][]telebot.InlineButton{}
	for i, trans := range transes {
		btn := telebot.InlineButton{
			Text:   fmt.Sprintf("%d. %s %d", i+1, trans.TypeTransaction(), trans.Value),
			Unique: fmt.Sprintf("%s:%d", CALLBACK_TRANS_STAT, trans.ID),
		}
		inlinekb = append(inlinekb, []telebot.InlineButton{btn})
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: inlinekb,
		ResizeKeyboard: true,
	}
}

func InlineTransDelete(trans models.MoneyTransaction) *telebot.ReplyMarkup {
	delete := telebot.InlineButton{
		Text:   TRANS_DELETE,
		Unique: fmt.Sprintf("%s:%d", CALLBACK_TRANS_DELETE, trans.ID),
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{delete},
		},
		ResizeKeyboard: true,
	}
}
