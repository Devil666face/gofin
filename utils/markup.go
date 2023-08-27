package utils

import (
	telebot "gopkg.in/telebot.v3"
)

var (
	Menu        = &telebot.ReplyMarkup{}
	NewTransBtn = telebot.ReplyButton{
		Text: TRANS_NEW,
	}
	RemoveMenu = &telebot.ReplyMarkup{RemoveKeyboard: true}
)

func init() {
	Menu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{NewTransBtn},
		},
		ResizeKeyboard: true,
	}
}
