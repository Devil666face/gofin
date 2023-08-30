package messages

import (
	"fmt"
	"time"
)

var (
	STATISTIC_TRANS_MONTH             = "Операции в текущем месяце"
	TRANS_DELETE                      = "Удалить"
	ERR_TRANS_DELETE                  = "Ошибка удаления операции"
	TRANS_SUCCESSFUL_DELELE           = "Операция успешно удалена"
	STATISTIC_CHANGE_TRANS_FOR_DELETE = "Для вывода подробностей или удаления - выберите операцию"
	STATISTIC_EXCEL_REPORT            = "Создать отчет Excel"
	ERR_STATISTIC_CREATE_REPORT       = "Произошла ошибка при создании отчета"
)

func StatInCurMonth() string {
	now := time.Now()
	return fmt.Sprintf("Операции %02d.%02d", now.Month(), now.Year())
}

func ErrGetTransForId(id uint) string {
	return fmt.Sprintf("Произошла ошибка при получении операции с id - %d", id)
}
