package database

import (
	"time"

	. "github.com/Devil666face/gofinabot/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db *gorm.DB
)

func Connect() error {
	db, err := gorm.Open(sqlite.Open(Cfg.Db), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	Db = db
	return nil
}

func Migrate(tables ...interface{}) error {
	err := Connect()
	if err != nil {
		return err
	}
	return Db.AutoMigrate(tables...)
}
