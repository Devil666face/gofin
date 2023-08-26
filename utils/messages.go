package utils

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
)

var (
	TRANS_NEW string = "Добавить операцию"
)

func ErrCreateUser(c telebot.Context) string {
	return fmt.Sprintf("Произошла ошибка при создании пользователя - %s", c.Chat().Username)
}

func SuccessfulCreateUser(c telebot.Context) string {
	return fmt.Sprintf("Пользователь %s - успешно создан", c.Chat().Username)
}

func UserAlreadyCreate(c telebot.Context) string {
	return fmt.Sprintf("Пользователь %s - уже создан", c.Chat().Username)
}
