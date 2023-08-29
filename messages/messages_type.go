package messages

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
)

var (
	TRANS_TYPE                 = "Категории операций"
	TYPE_ADD                   = "Добавить категорию для операций"
	CHANGE_TYPE_FOR_UPDATE     = "Выберите категорию для изменения"
	ADD_NEW_TYPE               = "Отправьте название категории"
	ERR_GET_ALL_TYPES_FOR_USER = "Произошла ошибка при получении всех категорий для вашего пользователя"
	OUR_TYPES                  = "Ваши категории:"
)

func SuccessfulCreateType(c telebot.Context) string {
	return fmt.Sprintf("Категория - %s - успешно создана", c.Message().Text)
}

func SuccessfulUpdateType(c telebot.Context) string {
	return fmt.Sprintf("Категория - %s - успешно обновлена", c.Message().Text)
}

func SuccessfulDeleteType(trtypename string) string {
	return fmt.Sprintf("Категория - %s - успешно удалена", trtypename)
}

func UpdateType(trtypename string) string {
	return fmt.Sprintf("Изменить категорию - %s", trtypename)
}

func UpdateTypeText(trtypename string) string {
	return fmt.Sprintf("Изменить название - %s", trtypename)
}

func DeleteTypeText(trtypename string) string {
	return fmt.Sprintf("Удалить - %s", trtypename)
}

func SendNewNameForType(trtypename string) string {
	return fmt.Sprintf("Введите новое имя для категории - %s", trtypename)
}

func ErrCreateType(c telebot.Context) string {
	return fmt.Sprintf("Произошла ошибка при создании категории - %s", c.Message().Text)
}

func ErrUpdateType(c telebot.Context) string {
	return fmt.Sprintf("Произошла ошибка при обновлении категории - %s", c.Message().Text)
}

func ErrGetTypeForId(id int64) string {
	return fmt.Sprintf("Произошла ошибка при получении категории с id - %d", id)
}

func ErrDeleteType(trtypename string) string {
	return fmt.Sprintf("Ошибка удаления категории - %s", trtypename)
}
