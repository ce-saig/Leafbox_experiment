package main

import (
	//"github.com/ce-saig/Leafbox_experiment"
	"github.com/ce-saig/Leafbox_experiment/model"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
	//	"github.com/jinzhu/gorm"
	//	"github.com/yosssi/ace"
	//"fmt"
	"net/http"
	//"strconv"
)

func ddHandler(w http.ResponseWriter, r *http.Request) {
	db := GetDB()
	var Items []model.Book
	db.Order("id desc").Limit(5).Find(&Items)
	data := map[string]interface{}{
		"Item": Items,
	}
	MakeView(w, "view/base", "view/home", data)
}

func ViewMediaHandler(w http.ResponseWriter, r *http.Request) {
	MakeView(w, "view/base", "view/book/media/braille", nil)
}
