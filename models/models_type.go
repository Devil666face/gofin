package models

import (
	. "github.com/Devil666face/gofinabot/store/database"

	"gorm.io/gorm"
)

type TypeTransaction struct {
	gorm.Model
	UserID            uint
	Type              string `gorm:"not null;index"`
	MoneyTransactions []MoneyTransaction
}

func (trtype *TypeTransaction) Create() error {
	if err := Db.Save(trtype); err != nil {
		return err.Error
	}
	return nil
}

func GetAllTypesForUser(id int64) ([]TypeTransaction, error) {
	var trtypes = []TypeTransaction{}
	if err := Db.Where("user_id = ?", id).Find(&trtypes); err != nil {
		return trtypes, err.Error
	}
	return trtypes, nil
}

func (trtype *TypeTransaction) Get(id int64) error {
	if err := Db.First(trtype, id); err != nil {
		return err.Error
	}
	return nil
}

func (trtype *TypeTransaction) Delete() error {
	if err := Db.Unscoped().Delete(trtype); err != nil {
		return err.Error
	}
	return nil
}

func (trtype *TypeTransaction) UpdateType(t string) error {
	trtype.Type = t
	if err := Db.Save(trtype); err != nil {
		return err.Error
	}
	return nil
}
