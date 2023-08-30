package database

import (
	"os"
	"path/filepath"
	"time"

	. "github.com/Devil666face/gofinabot/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db *gorm.DB
)

func setPath(file string) (string, error) {
	base, err := os.Getwd()
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(filepath.Join(base, file))
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(abs)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	return abs, nil
}

func Connect() error {
	path, err := setPath(Cfg.Db)
	if err != nil {
		return err
	}
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
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
