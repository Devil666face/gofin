package utils

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
)

var (
	TRANS_NEW                  = "Добавить операцию"
	TRANS_TYPE                 = "Категории операций"
	TYPE_ADD                   = "Добавить категорию для операций"
	BACK                       = "Назад"
	GO_BACK                    = "Возвращаемся назад"
	CONFIRM_USER               = "Добавить"
	IGNORE_USER                = "Игнорировать"
	CHANGE_TYPE_FOR_UPDATE     = "Выберите категорию для изменения"
	ADD_NEW_TYPE               = "Отправьте название категории"
	ERR_GET_ALL_TYPES_FOR_USER = "Произошла ошибка при получении всех категорий для вашего пользователя"
	OUR_TYPES                  = "Ваши категории"
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

func ErrCreateType(c telebot.Context) string {
	return fmt.Sprintf("Произошла ошибка при создании категории - %s", c.Message().Text)
}

func ErrUpdateType(c telebot.Context) string {
	return fmt.Sprintf("Произошла ошибка при обновлении категории - %s", c.Message().Text)
}

func SuccessfulCreateType(c telebot.Context) string {
	return fmt.Sprintf("Категория - %s - успешно создана", c.Message().Text)
}

func SuccessfulUpdateType(c telebot.Context) string {
	return fmt.Sprintf("Категория - %s - успешно обновлена", c.Message().Text)
}

func UpdateType(trtypename string) string {
	return fmt.Sprintf("Изменить категорию - %s", trtypename)
}

func ErrGetTypeForId(id int64) string {
	return fmt.Sprintf("Произошла ошибка при получении категории с id - %d", id)
}

func UpdateTypeText(trtypename string) string {
	return fmt.Sprintf("Изменить название - %s", trtypename)
}

func SendNewNameForType(trtypename string) string {
	return fmt.Sprintf("Введите новое имя для категории - %s", trtypename)
}

func DeleteTypeText(trtypename string) string {
	return fmt.Sprintf("Удалить - %s", trtypename)
}

func SuccessfulDeleteType(trtypename string) string {
	return fmt.Sprintf("Категория - %s - успешно удалена", trtypename)
}

func ErrDeleteType(trtypename string) string {
	return fmt.Sprintf("Ошибка удаления категории - %s", trtypename)
}
