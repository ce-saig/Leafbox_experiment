package main

import (
	//	"./model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/yosssi/ace"
	"net/http"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("view/base", "view/book/add", &ace.Options{DynamicReload: true})
	if err != nil {
		panic(err)
	}

	if err := tpl.Execute(w, nil); err != nil {
		panic(err)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	//--------ORM------
	db, err := gorm.Open("mysql", "root:web@tcp(192.168.56.102:3306)/leafbox?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
		return
	}

	var retData struct {
		Items []model.Book
	}

	db.Order("id desc").Limit(5).Find(&retData.Items)
	//------TPL-------
	data := map[string]interface{}{
		"Msg": []string{
			"Hello",
			"Ace",
		},
		"Item": retData.Items,
	}

	tpl, err := ace.Load("view/base", "view/home", &ace.Options{DynamicReload: true})
	if err != nil {
		panic(err)
	}

	if err := tpl.Execute(w, data); err != nil {
		panic(err)
	}
}

func GetRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/add", AddBookHandler)
	r.PathPrefix("/").Handler(
		http.StripPrefix("", http.FileServer(http.Dir("public/"))))
	return r
}
