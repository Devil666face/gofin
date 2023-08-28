package handlers

import (
	"log"

	. "github.com/Devil666face/gofinabot/database"
	. "github.com/Devil666face/gofinabot/models"
	"github.com/Devil666face/gofinabot/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

var (
	InputType              = fsm.NewStateGroup("type")
	InputTypeName          = InputType.New("type:name")
	InputTypeNameForUpdate = InputType.New("type:name_update")
)

var (
	TYPEID = "typeid"
)

func OnTransBtn(c telebot.Context) error {
	trtypes, err := GetAllTypesForUser(c.Chat().ID)
	if err != nil {
		return c.Send(utils.ERR_GET_ALL_TYPES_FOR_USER)
	}
	err = c.Send(utils.CHANGE_TYPE_FOR_UPDATE, utils.TypeMenu)
	if err != nil {
		log.Print(err)
	}
	return c.Send(utils.OUR_TYPES, utils.InlineTypes(trtypes))
}

func OnTypeAddBtn(c telebot.Context, s fsm.Context) error {
	s.Set(InputTypeName)
	return c.Send(utils.ADD_NEW_TYPE)
}

func OnTypeNameRecive(c telebot.Context, s fsm.Context) error {
	go s.Finish(true)
	trtype := TypeTransaction{
		Type:   c.Message().Text,
		UserID: uint(c.Chat().ID),
	}
	err := trtype.Create()
	if err != nil {
		return c.Send(utils.ErrCreateType(c))
	}
	return c.Send(utils.SuccessfulCreateType(c), utils.Menu)
}

func OnTypeNameForUpdateRecive(c telebot.Context, s fsm.Context) error {
	var (
		typeid int64
		t      TypeTransaction
	)
	val, ok := GetFromStore(c.Chat().ID, TYPEID)
	if !ok {
		log.Print("Not found value in store")
		return nil
	}
	typeid = val.(int64)
	err := t.Get(typeid)
	if err != nil {
		return c.Send(utils.ErrGetTypeForId(typeid))
	}
	err = t.UpdateType(c.Message().Text)
	if err != nil {
		return c.Send(utils.ErrUpdateType)
	}
	go s.Finish(true)
	return c.Send(utils.SuccessfulUpdateType(c), utils.Menu)
}

func getType(c telebot.Context) (TypeTransaction, error) {
	t := TypeTransaction{}
	typeid := utils.ToInt64(c.Get(CALLBACK_VAL))
	err := t.Get(typeid)
	if err != nil {
		return t, c.Send(utils.ErrGetTypeForId(typeid))
	}
	return t, nil
}

func OnUpdateTypeInlineBtn(c telebot.Context) error {
	delete(c)
	t, _ := getType(c)
	if c.Chat().ID != int64(t.UserID) {
		return nil
	}
	return c.Send(utils.UpdateType(t.Type), utils.UpdateTypeInline(t))
}

func OnUpdateTypeNameBtn(c telebot.Context, s fsm.Context) error {
	delete(c)
	t, _ := getType(c)
	if c.Chat().ID != int64(t.UserID) {
		return nil
	}
	go s.Set(InputTypeNameForUpdate)
	SetInStore(c.Chat().ID, TYPEID, int64(t.ID))
	return c.Send(utils.SendNewNameForType(t.Type))
}

func OnDeleteTypeBtn(c telebot.Context) error {
	delete(c)
	t, _ := getType(c)
	if c.Chat().ID != int64(t.UserID) {
		return nil
	}
	err := t.Delete()
	if err != nil {
		return c.Send(utils.ErrDeleteType(t.Type), utils.Menu)
	}
	return c.Send(utils.SuccessfulDeleteType(t.Type), utils.Menu)
}
