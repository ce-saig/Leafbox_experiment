package main

import (
	"./model"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/yosssi/ace"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	//response.Header().Set("Content-type", "text/html")
	//fmt.Fprintf(response, "Hey there!\n")

	//--------ORM------
	db, err := gorm.Open("mysql", "root:web@tcp(192.168.56.102:3306)/leafbox?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
		return
	}

	var retData struct {
		Items []model.Book
	}

	db.Find(&retData.Items)

	log.Println("All rows:")
	/*	for x, p := range retData.Items {
			log.Printf("    %d: %v\n", x, p.Title)
		}
	*/

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
