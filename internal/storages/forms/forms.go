package forms

import (
	"errors"

	"github.com/nlnaa11/BirthdayHelper/internal/model/updates"
)

type Forms struct {
	// key: name of dbTable, value: link
	forms map[string]string
	// key: name of dbTable, value: addedUsers
	added map[string]updates.User
}

func New(path, name, link string) (*Forms, error) {
	if path == "" || link == "" {
		return nil, errors.New("an empty path")
	}

	fm := Forms{}

	return &fm, nil
}

func (tg *Forms) Extract() error {
	// как организовать извлечение новых анкет
	return nil
}

func (tg *Forms) Save(db updates.Storage) error {
	// как оранизовать сохранение
	return nil
}
