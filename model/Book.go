package model

import (
	"time"
)

type Book struct {
	// db tag lets you specify the column name if it differs from the struct field

	Id            int
	Isbn          string `sql:"size:13"`
	Title         string `sql:"not null" `
	Author        string `sql:"size:50"`
	Translate     string `sql:"size:50"`
	Grade         string `sql:"size:50"`
	Bm_status     string `sql:"size:50"`
	Bm_note       string `sql:"size:50"`
	Bm_date       time.Time
	Setcs_status  string `sql:"size:50"`
	Setcs_note    string `sql:"size:50"`
	Setcs_date    time.Time
	Setds_status  string `sql:"size:50"`
	Setds_note    string `sql:"size:50"`
	Setds_date    time.Time
	Setcd_status  string `sql:"size:50"`
	Setcd_note    string `sql:"size:50"`
	Setcd_date    time.Time
	Setdvd_status string `sql:"size:50"`
	Setdvd_note   string `sql:"size:50"`
	Setdvd_date   time.Time
	Abstract      string
	Book_type     string `sql:"size:20"`
	Produce_no    string `sql:"size:20"`
	Original_no   string `sql:"size:20"`
	Pub_year      int
	Pub_no        int
	Publisher     string    `sql:"size:50"`
	Updated_at    time.Time `sql:"not null"`
	Created_at    time.Time `sql:"not null"`
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
