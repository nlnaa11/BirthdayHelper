package updates

import (
	"time"
)

type Sender interface {
	Extract() ([]Update, error)
}

type Receiver interface {
	Save([]Update) error
}

type ReceiverSender interface {
	Sender
	Receiver
}

type User struct {
	UserName string    `sql:"UserName"`
	Name     string    `sql:"Name"`
	Birthday time.Time `sql:"Birthday"`
}

type Type uint8

const (
	Add Type = 0
	Remove
	Birthday
)

type Update struct {
	user       *User
	typeUpdate Type
}
