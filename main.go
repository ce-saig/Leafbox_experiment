package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yosssi/ace"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	//response.Header().Set("Content-type", "text/html")
	//fmt.Fprintf(response, "Hey there!\n")

	tpl, err := ace.Load("view/base", "view/inner", &ace.Options{DynamicReload: true})
	if err != nil {
		panic(err)
	}
	if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
		panic(err)
	}
}

func main() {
	var host = flag.String("host", "127.0.0.1", "IP of host to run webserver on")
	var port = flag.Int("port", 8080, "Port to run webserver on")
	var staticPath = flag.String("staticPath", "public/", "Path to static files")

	flag.Parse()
	fmt.Println("Start")
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.PathPrefix("/").Handler(http.StripPrefix("", http.FileServer(http.Dir(*staticPath))))

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening on %s", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
