package model

import (
	"time"
)

type Braille struct {
	Id           int
	Book_id      int       `sql:"not null"`
	Reserved     bool      `sql:"not null"`
	Examiner     string    `sql:"size:50"`
	Produce_date time.Time `sql:"not null"`
	Pages        int       `sql:"not null"`
}

func (b Braille) TableName() string {
	return "braille"
}
