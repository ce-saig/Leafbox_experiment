package model

import (
	"time"
)

type Cd struct {
	Id           int
	BookId       int       `sql:"not null"`
	NumPart      int       `sql:"not null"`
	Reserved     bool      `sql:"not null"`
	ProducedDate time.Time `sql:"not null"`
	Notes        string    `sql:"not null"`
}

func (c Cd) TableName() string {
	return "cd"
}
