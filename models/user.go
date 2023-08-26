package models

import (
	"errors"

	"github.com/Devil666face/gofinabot/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TGID              uint   `gorm:"uniqueIndex;not null;index"`
	Username          string `gorm:"not null;default:noname;index"`
	IsAdmin           bool   `gorm:"default:false"`
	MoneyTransactions []MoneyTransaction
}

func (user *User) GetUserByTgID(id int64) error {
	err := database.DB.Where("tg_id = ?", id).Take(user)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return err.Error
	}
	return nil
}

func (user *User) Create() error {
	err := database.DB.Save(user)
	if err != nil {
		return err.Error
	}
	return nil
}
