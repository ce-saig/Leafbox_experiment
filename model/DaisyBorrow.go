package model

import (
	"time"
)

type DaisyBorrow struct {
	Id           int
	MemberId     int       `sql:"not null"`
	DaisyID      int       `sql:"not null"`
	BorrowedDate time.Time `sql:"not null"`
	ReturnedDate time.Time `sql:"not null"`
}

func (db DaisyBorrow) TableName() string {
	return "daisyborrow"
}
