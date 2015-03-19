package main

import (
	//"github.com/ce-saig/Leafbox_experiment"
	"github.com/ce-saig/Leafbox_experiment/model"
	_ "github.com/go-sql-driver/mysql"
	//	"github.com/jinzhu/gorm"
	//	"github.com/yosssi/ace"
	"fmt"
	"net/http"
	"strconv"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	MakeView(w, "view/base", "view/book/add", nil)
}

func AddBookPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()

	db := GetDB()

	pubNo, _ := strconv.Atoi(r.FormValue("pub_no"))
	pubYear, _ := strconv.Atoi(r.FormValue("pub_year"))

	newBook := model.Book{
		Isbn:        r.FormValue("isbn"),
		Title:       r.FormValue("title"),
		Author:      r.FormValue("author"),
		Translate:   r.FormValue("translate"),
		Publisher:   r.FormValue("publisher"),
		Pub_no:      pubNo,
		Pub_year:    pubYear,
		Produce_no:  r.FormValue("produce_no"),
		Original_no: r.FormValue("original_no"),
		Book_type:   r.FormValue("book_type"),
		Grade:       r.FormValue("grade"),
		Abstract:    r.FormValue("abstract"),
	}

	db.Create(&newBook)

	MakeView(w, "view/base", "view/book/add", nil) // Redirect
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	//--------ORM------
	db := GetDB()

	var Items []model.Book

	db.Order("id desc").Limit(5).Find(&Items)
	//------TPL-------
	data := map[string]interface{}{
		"Item": Items,
	}

	MakeView(w, "view/base", "view/home", data)

}
