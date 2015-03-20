package model

import (
	"time"
)

type Book struct {
	// db tag lets you specify the column name if it differs from the struct field

	Id           int
	Isbn         string `sql:"size:13"`
	Title        string `sql:"not null" `
	Author       string `sql:"size:50"`
	Translate    string `sql:"size:50"`
	Grade        string `sql:"size:50"`
	BmStatus     string `sql:"size:50"`
	BmNote       string `sql:"size:50"`
	BmDate       time.Time
	SetCsStatus  string `sql:"size:50"`
	SetCsNote    string `sql:"size:50"`
	SetCsDate    time.Time
	SetDsStatus  string `sql:"size:50"`
	SetDsNote    string `sql:"size:50"`
	SetDsDate    time.Time
	SetCdStatus  string `sql:"size:50"`
	SetCdNote    string `sql:"size:50"`
	SetCdDate    time.Time
	SetDvdStatus string `sql:"size:50"`
	SetDvdNote   string `sql:"size:50"`
	SetDvdDate   time.Time
	Abstract     string
	BookType     string `sql:"size:20"`
	ProduceNo    string `sql:"size:20"`
	OriginalNo   string `sql:"size:20"`
	PubYear      int
	PubNo        int
	Publisher    string    `sql:"size:50"`
	UpdatedAt    time.Time `sql:"not null"`
	CreatedAt    time.Time `sql:"not null"`
}

/*
func NewBook(title, author string) Book {
	return Book{
		Title:      title,
		Author:     author,
		Updated_at: time.Now().UnixNano(),
		Created_at: time.Now().UnixNano(),
	}
}
*/

func (b Book) TableName() string {
	return "book"
}
