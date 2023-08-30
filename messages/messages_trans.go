package messages

import (
	"fmt"

	"github.com/Devil666face/gofinabot/models"

	telebot "gopkg.in/telebot.v3"
)

var (
	EXPENSE                        = "Расход"
	INCOME                         = "Доход"
	TRANS_NEW                      = "Добавить операцию"
	GO_TRANS_NEW                   = "Создание операции"
	CHANGE_TRANS_BALANCE           = "Выберите тип опреации"
	CHANGE_TRANS_TYPE              = "Выберите категорию операции"
	SEND_TRANS_COMMENT             = "Отправьте комментарий для операции"
	EMPTY_COMMENT                  = "Оставить пустым"
	SEND_VALUE                     = "Отправьте сумму операции"
	TRANS_CREATE                   = "Создать"
	TRANS_CANCEL                   = "Отменить"
	TRANS_CANCEL_MESSAGE           = "Создание операции отменено"
	TRANS_SUCCESSFUL_CREATE        = "Операция успешно создана"
	ERR_TRANS_CREATE               = "Ошибка создания операции"
	ERR_INCOME_AND_NEGATIVE_VALUE  = "Значение дохода не может быть отрицательным"
	ERR_EXPENSE_AND_POSITIVE_VALUE = "Значение расхода не может быть положительным"
)

func CheckCreatedTrans(tr models.MoneyTransaction) string {
	return "Создать операцию со значениями:" + EnumerateTransFields(tr)
}

func EnumerateTransFields(tr models.MoneyTransaction) string {
	return fmt.Sprintf("\nТип: <b>%s</b>\nКатегория: <b>%s</b>\nСумма: <b>%d</b>\nКоментарий: <b>%s</b>\n", tr.Balance(), tr.TypeTransaction(), tr.Value, tr.Comment)
}

func ErrDisableValue(c telebot.Context) string {
	return fmt.Sprintf("Значение - %s - невозможно преобразовать в число, отправьте сумму операции еще раз", c.Message().Text)
}
