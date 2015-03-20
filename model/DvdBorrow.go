package model

import (
	"time"
)

type DvdBorrow struct {
	Id           int
	MemberId     int       `sql:"not null"`
	DvdID        int       `sql:"not null"`
	BorrowedDate time.Time `sql:"not null"`
	ReturnedDate time.Time `sql:"not null"`
}

func (db DvdBorrow) TableName() string {
	return "dvdborrow"
}
