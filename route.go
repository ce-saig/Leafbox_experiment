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

	rBookMenu.HandleFunc("/{id:[0-9]+}/edit", AddBookHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/edit", AddBookHandler).Methods("POST")

	rBookMenu.HandleFunc("/{id:[0-9]+}/braille/{id:[0-9]+}", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/braille/{id:[0-9]+}/edit", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/braille/add", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/braille/deleteAll", ViewMediaHandler)

	rBookMenu.HandleFunc("/{id:[0-9]+}/cassette/{id:[0-9]+}", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/cassette/{id:[0-9]+}/edit", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/cassette/add", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/cassette/deleteAll", ViewMediaHandler)

	rBookMenu.HandleFunc("/{id:[0-9]+}/daisy/{id:[0-9]+}", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/daisy/{id:[0-9]+}/edit", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/daisy/add", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/daisy/deleteAll", ViewMediaHandler)

	rBookMenu.HandleFunc("/{id:[0-9]+}/cd/{id:[0-9]+}", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/cd/{id:[0-9]+}/edit", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/cd/add", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/cd/deleteAll", ViewMediaHandler)

	rBookMenu.HandleFunc("/{id:[0-9]+}/dvd/{id:[0-9]+}", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/dvd/{id:[0-9]+}/edit", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/dvd/add", ViewMediaHandler)
	rBookMenu.HandleFunc("/{id:[0-9]+}/dvd/deleteAll", ViewMediaHandler)

	r.PathPrefix("/").Handler(
		http.StripPrefix("", http.FileServer(http.Dir("public/"))))
	return r
}
