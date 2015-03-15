package main

import (
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
	//_ "github.com/mattn/go-sqlite3"
	"./model"
	"github.com/go-gorp/gorp"
	"log"
)

func main() {
	// initialize the DbMap
	dbmap := initDb()
	defer dbmap.Db.Close()
	/*
		// delete any existing rows
		//err := dbmap.TruncateTables()
		//checkErr(err, "TruncateTables failed")
	*/
	/*
		// create two posts
		p1 := NewPost("Go 1.1 released!", "Lorem ipsum lorem ipsum")
		p2 := NewPost("Go 1.2 released!", "Lorem ipsum lorem ipsum")
	*/

	/*
		// insert rows - auto increment PKs will be set properly after the insert
		err = dbmap.Insert(&p1, &p2)
		checkErr(err, "Insert failed")
	*/

	/*
		// use convenience SelectInt
		count, err := dbmap.SelectInt("select count(*) from posts")
		checkErr(err, "select count(*) failed")
		log.Println("Rows after inserting:", count)
	*/
	/*
		// update a row
		p2.Title = "Go 1.2 is better than ever"
		count, err = dbmap.Update(&p2)
		checkErr(err, "Update failed")
		log.Println("Rows updated:", count)
	*/

	/*
		// fetch one row - note use of "post_id" instead of "Id" since column is aliased
		//
		// Postgres users should use $1 instead of ? placeholders
		// See 'Known Issues' below
		//
		err = dbmap.SelectOne(&p2, "select * from posts where post_id=?", p2.Id)
		checkErr(err, "SelectOne failed")
		log.Println("p2 row:", p2)
	*/

	// fetch all rows
	var books []model.Book
	_, err := dbmap.Select(&books, "select * from book ") //order by post_id
	checkErr(err, "Select failed")
	log.Println("All rows:")
	for x, p := range books {
		log.Printf("    %d: %v\n", x, p)
	}

	/*
		// delete row by PK
		count, err = dbmap.Delete(&p1)
		checkErr(err, "Delete failed")
		log.Println("Rows deleted:", count)
	*/

	/*	// delete row manually via Exec
		_, err = dbmap.Exec("delete from posts where post_id=?", 2)
		checkErr(err, "Exec failed")
	*/

	/*
		// confirm count is zero
		count, err = dbmap.SelectInt("select count(*) from posts")
		checkErr(err, "select count(*) failed")
		log.Println("Row count - should be zero:", count)
	*/

	log.Println("Done!")

}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	//db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	db, err := sql.Open("mymysql", "tcp:192.168.56.102:3306*leafbox/root/web")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	//dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(model.Book{}, "book").SetKeys(true, "Id")

	/*
		// create the table. in a production system you'd generally
		// use a migration tool, or create the tables via scripts
		err = dbmap.CreateTablesIfNotExists()
		checkErr(err, "Create tables failed")
	*/
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
