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

func GetAllTypesForUser(id int64) ([]TypeTransaction, error) {
	var trtypes = []TypeTransaction{}
	err := database.DB.Where("user_id = ?", id).Find(&trtypes)
	if err != nil {
		return trtypes, err.Error
	}
	return trtypes, nil
}

func (trtype *TypeTransaction) Get(id int64) error {
	err := database.DB.First(trtype, id)
	if err != nil {
		return err.Error
	}
	return nil
}

func (trtype *TypeTransaction) Delete() error {
	err := database.DB.Unscoped().Delete(trtype)
	if err != nil {
		return err.Error
	}
	return nil
}

func (trtype *TypeTransaction) UpdateType(t string) error {
	trtype.Type = t
	err := database.DB.Save(trtype)
	if err != nil {
		return err.Error
	}
	return nil
}
