package model

import (
	"time"
)

type CassetteBorrow struct {
	Id           int
	MemberId     int       `sql:"not null"`
	CassetteID   int       `sql:"not null"`
	BorrowedDate time.Time `sql:"not null"`
	ReturnedDate time.Time `sql:"not null"`
}

func (bb CassetteBorrow) TableName() string {
	return "cassetteborrow"
}
