package models

import (
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
