package model

import (
	"time"
)

type CdBorrow struct {
	Id           int
	MemberId     int       `sql:"not null"`
	CdID         int       `sql:"not null"`
	BorrowedDate time.Time `sql:"not null"`
	ReturnedDate time.Time `sql:"not null"`
}

func (bb CdBorrow) TableName() string {
	return "cdborrow"
}
