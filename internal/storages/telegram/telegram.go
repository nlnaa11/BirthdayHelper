package telegram

import (
	"bufio"
	"os"
	"strings"

	"github.com/nlnaa11/BirthdayHelper/internal/updates"
	"github.com/pkg/errors"
)

type Telegram struct {
	// value: name of dbTable + transferFile
	transfers []string
	// key: name of transferFile, value: removerUsers
	removedUsers map[string]updates.User
}

func New(path, name string) (*Telegram, error) {
	if path == "" || name == "" {
		return nil, errors.New("an empty path")
	}

	tg := Telegram{}

	return &tg, nil
}

func (tg *Telegram) Extract() ([]updates.Update, error) {
	// как организовать извлечение удаленных
	// извлекаем во внутренние данные
	return nil, nil
}

func (tg *Telegram) Save(upds []updates.Update) error {
	// как оранизовать сохранение обновлений (удаление)
	// опираясь на внутренние данные
	return nil
}

func (tg *Telegram) GetFormsLink() string {
	f, err := os.OpenFile(tg.transferFile, os.O_RDONLY, 0622)
	if err != nil {
		return ""
	}

	buffer := bufio.NewReader(f)
	chatInfo, err := buffer.ReadString('\n')
	if err != nil || chatInfo == "" {
		return ""
	}

	formsInfo := strings.Split(chatInfo, ",")[1]
	link := strings.SplitAfter(formsInfo, ":")[1]
	link = strings.TrimSuffix(strings.TrimPrefix(link, "\""), "\"")

	return link
}
