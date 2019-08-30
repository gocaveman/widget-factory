package main

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"
)

type DBConnString string

func NewDBConnString() DBConnString { return DBConnString(viper.GetString("db-conn")) }

type DBDriverName string

func NewDBDriverName() DBDriverName { return DBDriverName(viper.GetString("db-driver")) }

func NewMainStuff(c DBConnString, d DBDriverName) (*MainStuff, error) {
	db, err := sql.Open(string(d), string(c))
	if err != nil {
		return nil, err
	}
	return &MainStuff{
		DB: db,
	}, nil
}

type MainStuff struct {
	DB *sql.DB
}

func main() {

	mainStuff, err := Setup()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("HERE: %#v", mainStuff)
}
