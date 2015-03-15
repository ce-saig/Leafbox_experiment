package model

import (
	"time"
)

type Book struct {
	// db tag lets you specify the column name if it differs from the struct field

	Id            int
	Isbn          string
	Title         string
	Author        string
	Translate     string
	Grade         string
	Bm_status     string
	Bm_note       string
	Bm_date       time.Time
	Setcs_status  string
	Setcs_note    string
	Setcs_date    time.Time
	Setds_status  string
	Setds_note    string
	Setds_date    time.Time
	Setcd_status  string
	Setcd_note    string
	Setcd_date    time.Time
	Setdvd_status string
	Setdvd_note   string
	Setdvd_date   time.Time
	Abstract      string
	Book_type     string
	Produce_no    string
	Original_no   string
	Pub_year      int
	Pub_no        int
	Publisher     string
	Updated_at    time.Time
	Created_at    time.Time
}

func NewBook(title, author string) Book {
	return Book{
		Title:  title,
		Author: author,
		//Updated_at: time.Now().UnixNano(),
		//Created_at: time.Now().UnixNano(),
	}
}

func (b Book) TableName() string {
	return "book"
}
