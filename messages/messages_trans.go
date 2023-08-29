package messages

import (
	"fmt"

	"github.com/Devil666face/gofinabot/models"

	telebot "gopkg.in/telebot.v3"
)

var (
	EXPENSE              = "Расход"
	INCOME               = "Доход"
	TRANS_NEW            = "Добавить операцию"
	GO_TRANS_NEW         = "Создание операции"
	CHANGE_TRANS_BALANCE = "Выберите тип опреации"
	CHANGE_TRANS_TYPE    = "Выберите категорию операции"
	SEND_TRANS_COMMENT   = "Отправьте комментарий для операции"
	EMPTY_COMMENT        = "Оставить пустым"
	SEND_VALUE           = "Отправьте сумму операции"
)

func CheckCreatedTrans(tr models.MoneyTransaction) string {
	return fmt.Sprintf("Создать операцию со значениями:\n%s\n%s\n%t\n%d\n%s\n", tr.User(), tr.TypeTransaction(), tr.MoneyBalance, tr.Value, tr.Comment)
}

func ErrDisableValue(c telebot.Context) string {
	return fmt.Sprintf("Значение - %s - невозможно преобразовать в число, отправьте сумму операции еще раз", c.Message().Text)
}
