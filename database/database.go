package database

import (
	"time"

	. "github.com/Devil666face/gofinabot/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var Userstore map[int64]map[string]interface{}

func init() {
	Userstore = make(map[int64]map[string]interface{})
}

func SetInStore(id int64, k string, v interface{}) {
	if _, ok := Userstore[id]; !ok {
		Userstore[id] = make(map[string]interface{})
	}
	Userstore[id][k] = v
}

func GetFromStore(id int64, k string) (interface{}, bool) {
	val, ok := Userstore[id][k]
	return val, ok
}

func Connect() error {
	db, err := gorm.Open(sqlite.Open(Cfg.Db), &gorm.Config{
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
