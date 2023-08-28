package utils

import (
	"fmt"

	"github.com/Devil666face/gofinabot/models"
	telebot "gopkg.in/telebot.v3"
)

var (
	Menu        = &telebot.ReplyMarkup{}
	NewTransBtn = telebot.ReplyButton{
		Text: TRANS_NEW,
	}
	TypeTransBtn = telebot.ReplyButton{
		Text: TRANS_TYPE,
	}
	BackButton = telebot.ReplyButton{
		Text: BACK,
	}
)

var (
	TypeMenu      = &telebot.ReplyMarkup{}
	TypeAddButton = telebot.ReplyButton{
		Text: TYPE_ADD,
	}
)

var (
	CALLBACK_CONFIRM_USER string = "confirm_user"
	CALLBACK_IGNORE_USER  string = "ignore_user"
)

var (
	CALLBACK_TYPE        string = "type"
	CALLBACK_TYPE_UPDATE string = "type_update"
	CALLBACK_TYPE_DELETE string = "type_delete"
)

func init() {
	Menu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{NewTransBtn},
			{TypeTransBtn},
		},
		ResizeKeyboard: true,
	}

	TypeMenu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{TypeAddButton},
			{BackButton},
		},
		ResizeKeyboard: true,
	}
}

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

func InlineTypes(trtypes []models.TypeTransaction) *telebot.ReplyMarkup {
	btns := [][]telebot.InlineButton{}
	for _, v := range trtypes {
		btn := telebot.InlineButton{Text: v.Type, Unique: fmt.Sprintf("%s:%d", CALLBACK_TYPE, v.ID)}
		btns = append(btns, []telebot.InlineButton{btn})
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: btns,
		ResizeKeyboard: true,
	}
}

func UpdateTypeInline(trtype models.TypeTransaction) *telebot.ReplyMarkup {
	update := telebot.InlineButton{Text: UpdateTypeText(trtype.Type), Unique: fmt.Sprintf("%s:%d", CALLBACK_TYPE_UPDATE, trtype.ID)}
	delete := telebot.InlineButton{Text: DeleteTypeText(trtype.Type), Unique: fmt.Sprintf("%s:%d", CALLBACK_TYPE_DELETE, trtype.ID)}
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{update, delete},
		},
		ResizeKeyboard: true,
	}
}
