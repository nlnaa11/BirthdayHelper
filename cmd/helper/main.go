package main

import (
	"log"
	"os"
	"strings"

	"github.com/nlnaa11/BirthdayHelper/internal/storages/database"
	"github.com/nlnaa11/BirthdayHelper/internal/storages/forms"
	"github.com/nlnaa11/BirthdayHelper/internal/storages/telegram"
	"github.com/nlnaa11/BirthdayHelper/internal/updates"
	"github.com/pkg/errors"
)

func main() {

	// ..
	// path := os.Arguments[1]

	path := ""

	db, err := database.New(path)
	if err != nil {
		log.Fatal(err)
	}

	model, err := updates.New(path, db)
}

func foo() {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open the data folder: ")
	}
	if len(files) == 0 {
		return nil, errors.New(path + " is an empty directory")
	}

	tgs := make(map[string]*telegram.Telegram, 1)
	fms := make(map[string]*forms.Forms, 1)

	for i, file := range files {
		name := file.Name()

		if strings.HasSuffix(name, ".json") &&
			strings.HasPrefix(files[i-1].Name(), name) {

			tg, err := telegram.New(path, name)
			if err != nil {
				return nil, err
			}
			tgs[name] = tg

			link := tg.GetFormsLink()
			fm, err := forms.New(path, name, link)
			if err != nil {
				return nil, err
			}
			fms[name] = fm
		}
	}

}
