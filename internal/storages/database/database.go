package database

import (
	"database/sql"
	"time"

	"github.com/nlnaa11/BirthdayHelper/internal/updates"
	"github.com/pkg/errors"
)

type Database struct {
	data      *sql.DB
	birthdays []updates.User
}

func New(path string) (*Database, error) {
	if path == "" {
		return nil, errors.New("an empty path to db")
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "opening the database: ")
	}

	database := Database{
		data: db,
	}

	return &database, nil
}

func (db *Database) Add([]updates.User) error {
	return nil
}

func (db *Database) Remove([]updates.User) error {
	return nil
}

func (db *Database) Birthdays(date time.Time) ([]updates.User, error) {
	return nil, nil
}

/* The ReceiverSender interface implementation */

func (db *Database) Extact() ([]updates.Update, error) {
	return nil, nil
}

func (db *Database) Save(upds []updates.Update) error {
	return nil
}
