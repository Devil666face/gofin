package messages

import (
	"fmt"
	// telebot "gopkg.in/telebot.v3"
)

var (
	BACK    = "Назад"
	GO_BACK = "Возвращаемся назад"
)

func ErrSendMessage(username string) string {
	return fmt.Sprintf("Ошибка отправки сообщения пользователю - @%s", username)
}
