package database

import (
	"time"

	"github.com/Devil666face/gofinabot/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	db, err := gorm.Open(sqlite.Open(config.DB), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func Migrate(tables ...interface{}) error {
	err := Connect()
	if err != nil {
		return err
	}
	return DB.AutoMigrate(tables...)
}
