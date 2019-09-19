package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type DBConnString string

func NewDBConnString() DBConnString { return "root:root@/db" }

type DBDriverName string

func NewDBDriverName() DBDriverName { return "mysql" }

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

func NewDBConn(connString DBConnString, name DBDriverName) *sql.DB {
	db, err := sql.Open(string(name), string(connString))
	if err != nil {
		panic(err)
	}
	return db
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

	tx, err := mainStuff.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	upSql := []string{
		`DROP TABLE IF EXISTS widget_tag;`,
		`DROP TABLE IF EXISTS widget;`,
		`DROP TABLE IF EXISTS tag;`,
		`CREATE TABLE IF NOT EXISTS widget (
		 	widget_id varchar(64) not null,
		 	name varchar(255) not null,
		 	description text default null,
			PRIMARY KEY (widget_id)
		);`,
		`CREATE TABLE IF NOT EXISTS tag (
			tag_id varchar(64) not null,
			name varchar(255) not null,
			slug varchar(255) not null,
			description text default null,
			PRIMARY KEY (tag_id)
		);`,
		`CREATE TABLE IF NOT EXISTS widget_tag (
			widget_tag_id int(11) NOT NULL AUTO_INCREMENT,
			widget_id varchar(64) not null,
			tag_id varchar(64) not null,
			PRIMARY KEY (widget_tag_id),
			CONSTRAINT fk_widget FOREIGN KEY (widget_id) REFERENCES widget(widget_id),
			CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tag(tag_id)
		);`,
	}

	for _, query := range upSql {
		_, err = tx.Exec(query)
		if err != nil {
			log.Fatal(err)
		}

	}

	log.Fatal(http.ListenAndServe(":8080", mainStuff.Router))
}
