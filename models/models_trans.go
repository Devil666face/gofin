package models

import (
	"log"

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

// func (t *MoneyTransaction) String() string {
// 	return ""
// }

// func (trans *MoneyTransaction) Mone() string {

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
