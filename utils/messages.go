package utils

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
)

var (
	TRANS_NEW    string = "Добавить операцию"
	TRANS_TYPE   string = "Категории операций"
	TYPE_ADD     string = "Добавить категорию для операций"
	BACK         string = "Назад"
	CONFIRM_USER string = "Добавить"
	IGNORE_USER  string = "Игнорировать"
)

func ErrCreateUser(c telebot.Context) string {
	return fmt.Sprintf("Произошла ошибка при создании пользователя - @%s", c.Chat().Username)
}

func SuccessfulCreateUser(c telebot.Context) string {
	return fmt.Sprintf("Администратору отправлен запрос на создание вашего пользователя - @%s", c.Chat().Username)
}

func UserAlreadyCreate(c telebot.Context) string {
	return fmt.Sprintf("Пользователь @%s - уже создан", c.Chat().Username)
}

func AskAdminsForAddUser(c telebot.Context) string {
	return fmt.Sprintf("Добавить пользователя @%s - %d?", c.Chat().Username, c.Chat().ID)
}

func UserNotFound(id int64) string {
	return fmt.Sprintf("Пользователь с id - %d не найден", id)
}

func ErrUserUpdate(username string) string {
	return fmt.Sprintf("Ошибка обновления пользователя @%s", username)
}

func SuccessfulUpdateUser(username string) string {
	return fmt.Sprintf("Пользователь @%s - успешно обновлен", username)
}

func ErrSendMessage(username string) string {
	return fmt.Sprintf("Ошибка отправки сообщения пользователю - @%s", username)
}

func PermissionsForUserAdded(username string) string {
	return fmt.Sprintf("Администратор добавил вас - @%s", username)
}

func Back() string {
	return "Возвращаемся назад"
}

func ChangeTypeForUpdate() string {
	return "Выберите категорию для изменения"
}

func AddNewType() string {
	return "Отправьте название категории"
}

func ErrCreateType(c telebot.Context) string {
	return fmt.Sprintf("Произошла ошибка при создании категории - %s", c.Message().Text)
}

func SuccessfulCreateType(c telebot.Context) string {
	return fmt.Sprintf("Категория - %s - успешно создана", c.Message().Text)
}
