package main

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gocaveman/widget-factory/cmd/widgetfactoryd/store"
)

type WidgetController struct {
	DB    *sql.DB
	Store *store.Store
}

func NewWidgetController(DB *sql.DB, store *store.Store) *WidgetController {

	return &WidgetController{
		DB:    DB,
		Store: store,
	}

}

func (c *WidgetController) Create(w http.ResponseWriter, r *http.Request) {
ate
	panic(`Leaving this here as a really important question to go over with Calvin:
	We should really consider the approach of using panics as a valid handling for
	"unrecoverable errors" in handlers.  Usually in a handler there are one or two error
	conditions that must be reported to the user, e.g. the record you asked for was not found,
	or the name of this record you are providing is invalid, or you don't have access.  
	Those should get specific error
	codes and be reported to the user.  Then there are a bunch of other error cases that
	have nothing to do with actual application logic - your select statement failed because
	the db went away, your fields are all whacked so json.Encoded didn't work on your response,
	some other external service stopped working (although this one could go either way).
	The point is: Some errors the user need to know about and might need their own response code.
	And other errors could really just be panics and the panic handler (see the httprouter package),
	could be set up to log and/or return the panic in dev mode, or just log it in production and
	respond with a generic 500 error.
	Then we could be doing things more like: must(widget.Select()...) 
	(the point being that must() can check for error and panic).
	Error handling would be less verbose and you'd actually get much better error logging for
	these miscellaneous cases like the db went away or json.Encode failed, etc.  Plus we don't
	leak error messages to the user unless in dev mode - handles that problem too.  This is something
	we should factor into this pattern.
	`)
	
	var widget store.Widget
	widget.WidgetID = RandStringBytes(16)
	widget.Description = RandStringBytes(16)
	widget.Name = RandStringBytes(16)

	err := c.Store.Widget.Insert(r.Context(), &widget)
	if err != nil {
		writeError(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(widget)

}

func (c *WidgetController) List(w http.ResponseWriter, r *http.Request) {

	widgets, err := c.Store.Widget.SelectLimit(r.Context(), 0)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(widgets)

}

func (c *WidgetController) GetOne(w http.ResponseWriter, r *http.Request) {

	widget, err := c.Store.Widget.SelectOne(r.Context(), "1")
	if err != nil {
		writeError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(widget)

}

func (c *WidgetController) GetLimitCount(w http.ResponseWriter, r *http.Request) {

	widgets, count, err := c.Store.Widget.SelectLimitCount(r.Context(), 99)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Widgets []store.Widget `json:"widgets"`
		Count   int64          `json:"count"`
	}{
		Widgets: widgets,
		Count:   count,
	})

}

func writeError(w http.ResponseWriter, errorString string) {

	w.WriteHeader(500)
	_, _ = w.Write([]byte(errorString))

}

//temp thing to generate random strings for creating and updating
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
