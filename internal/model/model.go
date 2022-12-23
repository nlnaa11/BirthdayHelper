package model

import (
	"github.com/nlnaa11/BirthdayHelper/internal/updates"
)

type Model struct {
	stor     updates.ReceiverSender
	telegram updates.ReceiverSender
	forms    updates.Sender
}

func New(fms updates.Sender, stor, tg updates.ReceiverSender) *Model {
	return &Model{
		stor:     stor,
		telegram: tg,
		forms:    fms,
	}
}

func (m *Model) extractAndSaveUpdates(getter updates.Sender, stor updates.Receiver) error {
	// Извлечь обновления (разные типы) из себя
	upds, err := getter.Extract()
	if err != nil {
		return err
	}
	// Сохранить обновления (добавить или удалить данные)
	err = stor.Save(upds)
	if err != nil {
		return err
	}

	return nil
}

/**
 * Update:
 * 1. extract UPDATES from forms + save to db
 * 2. extract UPDATES from transfer + remove from db  (clear transfer)
 * 3. select(=extract) Birthdays from db + save to transfer (fill transfer)
 */
// эта функция вызывается из мэйн (?) в 00:00
func (m *Model) Update() error {
	// 1. Extract Upds from forms and save to db
	if err := m.extractAndSaveUpdates(m.forms, m.stor); err != nil {
		return err
	}

	// 2. Extract Upds from transfers and remove from db
	if err := m.extractAndSaveUpdates(m.telegram, m.stor); err != nil {
		return err
	}

	// 3. Extact Births from db and save to transers
	if err := m.extractAndSaveUpdates(m.stor, m.telegram); err != nil {
		return err
	}

	return nil
}

// эта функция вызывается из мэйн (?) в 8:58
