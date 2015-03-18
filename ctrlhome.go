package main

import (
	//"github.com/ce-saig/Leafbox_experiment"
	"github.com/ce-saig/Leafbox_experiment/model"
	_ "github.com/go-sql-driver/mysql"
	//	"github.com/jinzhu/gorm"
	//	"github.com/yosssi/ace"
	"fmt"
	"net/http"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	MakeView(w, "view/base", "view/book/add", nil)
}

func AddBookPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("title"))
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
