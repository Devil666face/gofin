package models

import (
	"github.com/Devil666face/gofinabot/database"

	"gorm.io/gorm"
)

type TypeTransaction struct {
	gorm.Model
	UserID            uint
	Type              string `gorm:"not null;index"`
	MoneyTransactions []MoneyTransaction
}

func (trtype *TypeTransaction) Create() error {
	err := database.DB.Save(trtype)
	if err != nil {
		return err.Error
	}
	return nil
}
