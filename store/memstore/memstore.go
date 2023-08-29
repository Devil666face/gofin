package memstore

import (
	"github.com/vitaliy-ukiru/fsm-telebot/storages/memory"
)

var (
	Memstore *memory.Storage
)

func Store() {
	store := memory.NewStorage()
	Memstore = store
}
