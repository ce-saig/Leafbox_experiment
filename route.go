package main

import (
	"github.com/ce-saig/Leafbox_experiment/model"
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
	db := getDB()

	var Items []model.Book

	db.Order("id desc").Limit(5).Find(&Items)
	//------TPL-------
	data := map[string]interface{}{
		"Item": Items,
	}

	makeView(w, "view/base", "view/home", data)

}

func makeView(w http.ResponseWriter, base, inner string, data interface{}) {
	tpl, err := ace.Load(base, inner, &ace.Options{DynamicReload: true})
	if err != nil {
		panic(err)
	}

	if err := tpl.Execute(w, data); err != nil {
		panic(err)
	}
}

func getDB() gorm.DB {

	config := getConfig()
	sql := config["DB_USER"].(string) + ":" + config["DB_PWD"].(string) + "@tcp(localhost:" + config["DB_PORT"].(string) + ")/leafbox?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sql)
	if err != nil {
		panic(err)
	}
	return db
}

func getRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/add", AddBookHandler)
	r.PathPrefix("/").Handler(
		http.StripPrefix("", http.FileServer(http.Dir("public/"))))
	return r
}
