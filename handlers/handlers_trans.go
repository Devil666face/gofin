package handlers

import (
	"log"

	kb "github.com/Devil666face/gofinabot/markup"
	m "github.com/Devil666face/gofinabot/messages"
	// . "github.com/Devil666face/gofinabot/models"
	// "github.com/Devil666face/gofinabot/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

var (
	InputTrans        = fsm.NewStateGroup("trans")
	InputTransBalance = InputTrans.New("balance")
	InputTransType    = InputTrans.New("type")
	InputTransComment = InputTrans.New("comment")
	InputTransValue   = InputTrans.New("value")
)

func OnAddTransBtn(c telebot.Context, s fsm.Context) error {
	if err := s.Set(InputTransBalance); err != nil {
		log.Print(err)
	}
	if err := c.Send(m.GO_TRANS_NEW, kb.TransMenu); err != nil {
		log.Print(err)
	}
	return c.Send(m.CHANGE_TRANS_BALANCE, kb.BalanceInline)
}
