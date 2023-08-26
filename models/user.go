package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TGID              uint   `gorm:"uniqueIndex;not null;index"`
	Username          string `gorm:"not null;default:noname;index"`
	IsAdmin           bool   `gorm:"default:false"`
	MoneyTransactions []MoneyTransaction
}
