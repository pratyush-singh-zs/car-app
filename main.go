package main

import (
	// "database/sql/driver"
	"log"
	"net/http"

	// "os"
	storecar "github.com/pratyush-singh-zs/car-app/datastore/car"
	"github.com/pratyush-singh-zs/car-app/driver"
	handlercar "github.com/pratyush-singh-zs/car-app/handler/car"
)

// var urlDsn = "root:cools10cj@tcp(127.0.0.1:3306)/mydb"

func main() {
	conf := driver.MySQLConfig{
		Host:     "127.0.0.1",
		User:     "root",
		Password: "cools10cj",
		Port:     "3306",
		Db:       "mydb",
	}

	var err error

	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		log.Println("could not connect to sql, err:", err)
		return
	}
	datastore := storecar.New(db)
	handler := handlercar.New(datastore)
	// handlercar.CarHandler
	http.HandleFunc("/car", handler.Hander)
	//createcar
	//getcar
	//updatecar
	log.Fatal(http.ListenAndServe(":8080", nil))
}
