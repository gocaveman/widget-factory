package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

type DBConnString string

func NewDBConnString() DBConnString { return DBConnString(viper.GetString("db-conn")) }

type DBDriverName string

func NewDBDriverName() DBDriverName { return DBDriverName(viper.GetString("db-driver")) }

func NewMainStuff(c DBConnString, d DBDriverName, router *httprouter.Router) (*MainStuff, error) {
	db, err := sql.Open(string(d), string(c))
	if err != nil {
		return nil, err
	}
	return &MainStuff{
		DB:     db,
		Router: router,
	}, nil
}

type MainStuff struct {
	DB     *sql.DB
	Router *httprouter.Router
}

func main() {

	mainStuff, err := Setup()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", mainStuff.Router))
}
