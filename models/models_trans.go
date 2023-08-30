package models

import (
	"log"

	. "github.com/Devil666face/gofinabot/store/database"
	"github.com/Devil666face/gofinabot/utils"

	"gorm.io/gorm"
)

type MoneyTransaction struct {
	gorm.Model
	UserID            uint
	TypeTransactionID uint
	MoneyBalance      bool `gorm:"default:false"`
	Value             int  `gorm:"default:0"`
	Comment           string
}

func (trans *MoneyTransaction) Get(id uint) error {
	if err := Db.First(trans, id); err != nil {
		return err.Error
	}
	return nil
}

func (trans *MoneyTransaction) Delete() error {
	if err := Db.Unscoped().Delete(trans); err != nil {
		return err.Error
	}
	return nil
}

func (trans *MoneyTransaction) Balance() string {
	if trans.MoneyBalance {
		return "Доход"
	}
	return "Расход"
}

func (trans *MoneyTransaction) Create() error {
	if err := Db.Save(trans); err != nil {
		return err.Error
	}
	return nil
}

func (trans *MoneyTransaction) TypeTransaction() TypeTransaction {
	var trtype TypeTransaction
	if err := trtype.Get(trans.TypeTransactionID); err != nil {
		log.Print(err)
	}
	return trtype
}

func (trans *MoneyTransaction) User() User {
	var user User
	if err := user.GetUserByTgID(int64(trans.UserID)); err != nil {
		log.Print(err)
	}
	return user
}

func GetAllTransInCurMonthForUser(userid int64) ([]MoneyTransaction, error) {
	start, end := utils.GetStartAndEndOfMonth()
	var trans = []MoneyTransaction{}
	if err := Db.Where("user_id = ?", uint(userid)).Where("created_at > ?", start).Where("created_at < ?", end).Find(&trans); err != nil {
		return trans, err.Error
	}
	return trans, nil
}
