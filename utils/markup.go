package utils

import (
	"fmt"

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
