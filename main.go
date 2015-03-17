package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var host = flag.String("host", "127.0.0.1", "IP of host to run webserver on")
	var port = flag.Int("port", 8080, "Port to run webserver on")

	flag.Parse()
	fmt.Println("Start")
	router := GetRoute()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening on %s", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

}
