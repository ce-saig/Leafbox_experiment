package main

/*
import (
	"./model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
)

var (
	db            gorm.DB
	sqlConnection string
)

func main() {
	var err error

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
	for x, p := range retData.Items {
		log.Printf("    %d: %v\n", x, p.Title)
	}

}
*/
