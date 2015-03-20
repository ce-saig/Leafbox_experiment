package model

import (
	"time"
)

type BrailleBorrow struct {
	Id           int
	MemberId     int       `sql:"not null"`
	BrailleID    int       `sql:"not null"`
	BorrowedDate time.Time `sql:"not null"`
	ReturnedDate time.Time `sql:"not null"`
}

func (bb BrailleBorrow) TableName() string {
	return "brailleborrow"
}
