package model

import (
	"time"
)

type Cassette struct {
	Id          int
	BookId      int       `sql:"not null"`
	NumPart     int       `sql:"not null"`
	Reserved    bool      `sql:"not null"`
	ProduceDate time.Time `sql:"not null"`
	Notes       string    `sql:"not null"`
}

func (c Cassette) TableName() string {
	return "cassette"
}
