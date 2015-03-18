package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yosssi/ace"
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

func MakeView(w http.ResponseWriter, base, inner string, data interface{}) {
	tpl, err := ace.Load(base, inner, &ace.Options{DynamicReload: true})
	if err != nil {
		panic(err)
	}

	if err := tpl.Execute(w, data); err != nil {
		panic(err)
	}
}

func GetDB() gorm.DB {

	config := getConfig()
	sql := config["DB_USER"].(string) + ":" + config["DB_PWD"].(string) + "@tcp(localhost:" + config["DB_PORT"].(string) + ")/leafbox?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sql)
	if err != nil {
		panic(err)
	}
	return db
}
