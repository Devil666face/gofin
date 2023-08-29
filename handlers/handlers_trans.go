package handlers

import (
	"fmt"
	"log"

	kb "github.com/Devil666face/gofinabot/markup"
	m "github.com/Devil666face/gofinabot/messages"
	. "github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils"

	// "github.com/Devil666face/gofinabot/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

var (
	InputTrans             = fsm.NewStateGroup("trans")
	InputTransBalanceState = InputTrans.New("balance")
	InputTransTypeState    = InputTrans.New("type")
	InputTransCommentState = InputTrans.New("comment")
	InputTransValueState   = InputTrans.New("value")
	InputTransCheckAll     = InputTrans.New("check")
)

func OnAddTransBtn(c telebot.Context, s fsm.Context) error {
	if err := s.Set(InputTransBalanceState); err != nil {
		log.Print(err)
	}
	if err := c.Send(m.GO_TRANS_NEW, kb.TransMenu); err != nil {
		log.Print(err)
	}
	return c.Send(m.CHANGE_TRANS_BALANCE, kb.BalanceInline)
}

func OnTransBalanceInlineBtn(c telebot.Context, s fsm.Context) error {
	delete(c)
	if err := s.Update(InputTransBalanceState.GoString(), utils.StoBoll(c.Get(CALLBACK_VAL))); err != nil {
		log.Print(err)
	}
	if err := s.Set(InputTransTypeState); err != nil {
		log.Print(err)
	}
	trtypes, err := GetAllTypesForUser(c.Chat().ID)
	if err != nil {
		return c.Send(m.ERR_GET_ALL_TYPES_FOR_USER)
	}

	return c.Send(m.CHANGE_TRANS_TYPE, kb.InlineTypesForAddTrans(trtypes))
}

func OnTransTypeInlibeBtn(c telebot.Context, s fsm.Context) error {
	delete(c)
	if err := s.Update(InputTransTypeState.GoString(), utils.ToUint(c.Get(CALLBACK_VAL))); err != nil {
		log.Print(err)
	}
	if err := s.Set(InputTransCommentState); err != nil {
		log.Print(err)
	}
	return c.Send(m.SEND_TRANS_COMMENT, kb.TransCommentInline)
}

func OnTransCommentRecive(c telebot.Context, s fsm.Context) error {
	delete(c)
	if err := c.Callback(); err != nil {
		// If callback - set empty
		if err := s.Update(InputTransCommentState.GoString(), ""); err != nil {
			log.Print(err)
		}
	} else {
		if err := s.Update(InputTransCommentState.GoString(), c.Message().Text); err != nil {
			log.Print(err)
		}
	}
	if err := s.Set(InputTransValueState); err != nil {
		log.Print(err)
	}
	return c.Send(m.SEND_VALUE)
}
func OnTransValueRecive(c telebot.Context, s fsm.Context) error {
	delete(c)
	var (
		balance  bool
		trtypeid uint
		comment  string
	)

	value, err := utils.ToInt(c.Message().Text)
	if err != nil {
		return c.Send(m.ErrDisableValue(c))
	}

	if err := s.Get(InputTransBalanceState.GoString(), &balance); err != nil {
		log.Print(err)
	}
	if err := s.Get(InputTransTypeState.GoString(), &trtypeid); err != nil {
		log.Print(err)
	}
	if err := s.Get(InputTransCommentState.GoString(), &comment); err != nil {
		log.Print(err)
	}

	tr := MoneyTransaction{
		UserID:            uint(c.Chat().ID),
		TypeTransactionID: trtypeid,
		MoneyBalance:      balance,
		Value:             value,
		Comment:           comment,
	}

	if err := s.Update(InputTransCheckAll.GoString(), tr); err != nil {
		log.Print(err)
	}

	if err := s.Set(InputTransCheckAll); err != nil {
		log.Print(err)
	}

	return c.Send(m.CheckCreatedTrans(tr))
}

func OnTransAddInlineBtn(c telebot.Context, s fsm.Context) error {
	if err := s.Finish(true); err != nil {
		log.Print(err)
	}
	return nil
}
