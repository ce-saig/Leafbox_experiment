package main

import (
	//"github.com/ce-saig/Leafbox_experiment/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//	"github.com/jinzhu/gorm"
	//	"github.com/yosssi/ace"
	"net/http"
)

func getRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/search", RootHandler)
	r.HandleFunc("/login", RootHandler)
	r.HandleFunc("/logout", RootHandler)

	rBook := r.PathPrefix("/book").Subrouter()
	rBook.HandleFunc("/add", AddBookHandler).Methods("GET")
	rBook.HandleFunc("/add", AddBookPostHandler).Methods("POST")
	/*
		r.HandleFunc("/add", AddBookHandler)
		r.HandleFunc("/add", AddBookHandler)
		r.HandleFunc("/add", AddBookHandler)
		r.HandleFunc("/add", AddBookHandler)
		r.HandleFunc("/add", AddBookHandler)
		r.HandleFunc("/add", AddBookHandler)
		r.HandleFunc("/add", AddBookHandler)
		r.HandleFunc("/add", AddBookHandler)
		r.HandleFunc("/add", AddBookHandler)
	*/
	r.PathPrefix("/").Handler(
		http.StripPrefix("", http.FileServer(http.Dir("public/"))))
	return r
}
