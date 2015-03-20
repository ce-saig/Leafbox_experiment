package model

import (
	"time"
)

type Dvd struct {
	Id          int
	BookId      int       `sql:"not null"`
	NumPart     int       `sql:"not null"`
	Reserved    bool      `sql:"not null"`
	ProduceDate time.Time `sql:"not null"`
	Notes       string    `sql:"not null"`
}

func (d Dvd) TableName() string {
	return "dvd"
}
