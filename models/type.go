package models

import (
	"gorm.io/gorm"
)

type TypeTransaction struct {
	gorm.Model
	Type              string `gorm:"not null;index"`
	MoneyTransactions []MoneyTransaction
}
