package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getConfig() map[string]interface{} {
	dat, _ := ioutil.ReadFile("app.json")
	var f interface{}
	err := json.Unmarshal([]byte(dat), &f)
	if err != nil {
		fmt.Println(err)
	}
	return f.(map[string]interface{})
}

func main() {
	var host = flag.String("host", "127.0.0.1", "IP of host to run webserver on")
	var port = flag.Int("port", 8080, "Port to run webserver on")

	flag.Parse()
	fmt.Println("Start")

	router := getRoute()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening on %s", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

}
