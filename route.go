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

	rBookMenu := rBook.PathPrefix("/").Subrouter()
	rBookMenu.HandleFunc("/{id:[0-9]+}", ViewBookHandler)
	/*
		rBookMenu.HandleFunc("/edit", AddBookHandler)
		rBookMenu.HandleFunc("/edit", AddBookHandler)

		rBookMenu.HandleFunc("/braille/{id:[0-9]+}", AddBookHandler)
		rBookMenu.HandleFunc("/braille/{id:[0-9]+}/edit", AddBookHandler)
		rBookMenu.HandleFunc("/braille/add", AddBookHandler)
		rBookMenu.HandleFunc("/braille/deleteAll", AddBookHandler)

		rBookMenu.HandleFunc("/cassette/{id:[0-9]+}", AddBookHandler)
		rBookMenu.HandleFunc("/cassette/{id:[0-9]+}/edit", AddBookHandler)
		rBookMenu.HandleFunc("/cassette/add", AddBookHandler)
		rBookMenu.HandleFunc("/cassette/deleteAll", AddBookHandler)

		rBookMenu.HandleFunc("/daisy/{id:[0-9]+}", AddBookHandler)
		rBookMenu.HandleFunc("/daisy/{id:[0-9]+}/edit", AddBookHandler)
		rBookMenu.HandleFunc("/daisy/add", AddBookHandler)
		rBookMenu.HandleFunc("/daisy/deleteAll", AddBookHandler)

		rBookMenu.HandleFunc("/cd/{id:[0-9]+}", AddBookHandler)
		rBookMenu.HandleFunc("/cd/{id:[0-9]+}/edit", AddBookHandler)
		rBookMenu.HandleFunc("/cd/add", AddBookHandler)
		rBookMenu.HandleFunc("/cd/deleteAll", AddBookHandler)

		rBookMenu.HandleFunc("/dvd/{id:[0-9]+}", AddBookHandler)
		rBookMenu.HandleFunc("/dvd/{id:[0-9]+}/edit", AddBookHandler)
		rBookMenu.HandleFunc("/dvd/add", AddBookHandler)
		rBookMenu.HandleFunc("/dvd/deleteAll", AddBookHandler)
	*/
	r.PathPrefix("/").Handler(
		http.StripPrefix("", http.FileServer(http.Dir("public/"))))
	return r
}
