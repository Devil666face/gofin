package models

import (
	"log"

	. "github.com/Devil666face/gofinabot/store/database"

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
