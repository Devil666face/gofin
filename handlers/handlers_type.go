package handlers

import (
	"log"

	kb "github.com/Devil666face/gofinabot/markup"
	m "github.com/Devil666face/gofinabot/messages"
	. "github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

var (
	InputTypeName               = fsm.NewStateGroup("type")
	InputTypeNameState          = InputTypeName.New("name")
	InputTypeNameForUpdateState = InputTypeName.New("name_update")
)

func OnTypeBtn(c telebot.Context) error {
	trtypes, err := GetAllTypesForUser(c.Chat().ID)
	if err != nil {
		return c.Send(m.ERR_GET_ALL_TYPES_FOR_USER)
	}
	if err = c.Send(m.CHANGE_TYPE_FOR_UPDATE, kb.TypeMenu); err != nil {
		log.Print(err)
	}
	return c.Send(m.OUR_TYPES, kb.InlineTypes(trtypes))
}

func OnTypeAddBtn(c telebot.Context, s fsm.Context) error {
	if err := s.Set(InputTypeNameState); err != nil {
		log.Print(err)
	}
	return c.Send(m.ADD_NEW_TYPE)
}

func OnTypeNameRecive(c telebot.Context, s fsm.Context) error {
	if err := s.Finish(true); err != nil {
		log.Print(err)
	}
	trtype := TypeTransaction{
		Type:   c.Message().Text,
		UserID: uint(c.Chat().ID),
	}
	if err := trtype.Create(); err != nil {
		return c.Send(m.ErrCreateType(c))
	}
	return c.Send(m.SuccessfulCreateType(c), kb.Menu)
}

func OnTypeNameForUpdateRecive(c telebot.Context, s fsm.Context) error {
	var (
		typeid uint
		t      TypeTransaction
	)
	if err := s.Get(InputTypeNameForUpdateState.GoString(), &typeid); err != nil {
		log.Print(err)
	}
	if err := t.Get(typeid); err != nil {
		return c.Send(m.ErrGetTypeForId(typeid))
	}
	if err := t.UpdateType(c.Message().Text); err != nil {
		return c.Send(m.ErrUpdateType)
	}
	if err := s.Finish(true); err != nil {
		log.Print(err)
	}
	return c.Send(m.SuccessfulUpdateType(c), kb.Menu)
}

func getType(c telebot.Context) (TypeTransaction, error) {
	t := TypeTransaction{}
	typeid := utils.ToUint(c.Get(CALLBACK_VAL))
	if err := t.Get(typeid); err != nil {
		return t, c.Send(m.ErrGetTypeForId(typeid))
	}
	if c.Chat().ID != int64(t.UserID) {
		return t, telebot.ErrNotFound
	}
	return t, nil
}

func OnUpdateCurrentTypeInlineBtn(c telebot.Context) error {
	delete(c)
	t, err := getType(c)
	if err != nil {
		return err
	}
	return c.Send(m.UpdateType(t.Type), kb.UpdateTypeInline(t))
}

func OnUpdateTypeNameInlineBtn(c telebot.Context, s fsm.Context) error {
	delete(c)
	t, err := getType(c)
	if err != nil {
		return err
	}
	if err := s.Set(InputTypeNameForUpdateState); err != nil {
		log.Print(err)
	}
	if err := s.Update(InputTypeNameForUpdateState.GoString(), t.ID); err != nil {
		log.Print(err)
	}
	return c.Send(m.SendNewNameForType(t.Type))
}

func OnDeleteTypeInlineBtn(c telebot.Context) error {
	delete(c)
	t, err := getType(c)
	if err != nil {
		return err
	}
	if err := t.Delete(); err != nil {
		return c.Send(m.ErrDeleteType(t.Type), kb.Menu)
	}
	return c.Send(m.SuccessfulDeleteType(t.Type), kb.Menu)
}
