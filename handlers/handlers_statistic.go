package handlers

import (
	"log"
	"os"

	kb "github.com/Devil666face/gofinabot/markup"
	m "github.com/Devil666face/gofinabot/messages"
	. "github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils"
	"github.com/Devil666face/gofinabot/utils/excel"

	// "github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

func OnMonthStatBtn(c telebot.Context) error {
	delete(c)
	trans, err := GetAllTransInCurMonthForUser(c.Chat().ID)
	if err != nil {
		log.Print(err)
	}
	if err := c.Send(m.StatInCurMonth(), kb.StatMenu); err != nil {
		log.Print(err)
	}
	return c.Send(m.STATISTIC_CHANGE_TRANS_FOR_DELETE, kb.InlineTransStatList(trans))
}

func getTrans(c telebot.Context) (MoneyTransaction, error) {
	t := MoneyTransaction{}
	trid := utils.ToUint(c.Get(CALLBACK_VAL))
	if err := t.Get(trid); err != nil {
		return t, c.Send(m.ErrGetTransForId(trid))
	}
	if c.Chat().ID != int64(t.UserID) {
		return t, telebot.ErrNotFound
	}
	return t, nil
}

func OnTransStatInlineBtn(c telebot.Context) error {
	delete(c)
	trans, err := getTrans(c)
	if err != nil {
		return err
	}
	return c.Send(m.EnumerateTransFields(trans), kb.InlineTransDelete(trans))
}

func OnTransDeleteInlineBtn(c telebot.Context) error {
	delete(c)
	trans, err := getTrans(c)
	if err != nil {
		return err
	}
	if err := trans.Delete(); err != nil {
		return c.Send(m.ERR_TRANS_DELETE, kb.Menu)
	}
	return c.Send(m.TRANS_SUCCESSFUL_DELELE, kb.Menu)
}

func OnExcelReportBtn(c telebot.Context) error {
	trans, err := GetAllTransInCurMonthForUser(c.Chat().ID)
	if err != nil {
		log.Print(err)
	}
	r := excel.New(trans)
	if err := r.Save(); err != nil {
		if err := c.Send(m.ERR_STATISTIC_CREATE_REPORT, kb.Menu); err != nil {
			log.Print(err)
		}
	}
	f := &telebot.Document{
		File:     telebot.FromDisk(r.FileName),
		FileName: r.FileName,
	}
	defer func() {
		if err := os.Remove(r.FileName); err != nil {
			log.Print(err)
		}
	}()
	return c.Send(f, kb.Menu)
}
