package main

import (
	//"github.com/ce-saig/Leafbox_experiment"
	"github.com/ce-saig/Leafbox_experiment/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//	"github.com/jinzhu/gorm"
	//	"github.com/yosssi/ace"
	"fmt"
	"net/http"
	"strconv"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	db := GetDB()
	var Items []model.Book
	db.Order("id desc").Limit(5).Find(&Items)
	data := map[string]interface{}{
		"Item": Items,
	}
	MakeView(w, "view/base", "view/home", data)
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	MakeView(w, "view/base", "view/book/add", nil)
}

func AddBookPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()

	db := GetDB()

	pubNo, _ := strconv.Atoi(r.FormValue("pub_no"))
	pubYear, _ := strconv.Atoi(r.FormValue("pub_year"))

	newBook := model.Book{
		Isbn:       r.FormValue("isbn"),
		Title:      r.FormValue("title"),
		Author:     r.FormValue("author"),
		Translate:  r.FormValue("translate"),
		Publisher:  r.FormValue("publisher"),
		PubNo:      pubNo,
		PubYear:    pubYear,
		ProduceNo:  r.FormValue("produce_no"),
		OriginalNo: r.FormValue("original_no"),
		BookType:   r.FormValue("book_type"),
		Grade:      r.FormValue("grade"),
		Abstract:   r.FormValue("abstract"),
	}

	db.Create(&newBook)

	MakeView(w, "view/base", "view/book/add", nil) // Redirect
}

func ViewBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db := GetDB()

	var Book []model.Book
	db.First(&Book, id)

	var Braille []model.Braille
	db.Find(&Braille)

	data := map[string]interface{}{
		"Book":    Book,
		"Braille": Braille,
	}
	MakeView(w, "view/base", "view/book/view", data)
}
