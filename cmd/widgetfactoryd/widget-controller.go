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

	var widget store.Widget
	widget.WidgetID = RandStringBytes(16)
	widget.Description = RandStringBytes(16)
	widget.Name = RandStringBytes(16)

	if err := c.Store.Widget().Insert(r.Context(), &widget); err != nil {
		writeError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(widget)

}

func (c *WidgetController) List(w http.ResponseWriter, r *http.Request) {

	widgets, err := c.Store.Widget().SelectLimit(r.Context(), 0)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(widgets)

}

func (c *WidgetController) GetOne(w http.ResponseWriter, r *http.Request) {

	widget, err := c.Store.Widget().SelectOne(r.Context(), "1")
	if err != nil {
		writeError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(widget)

}

func (c *WidgetController) GetLimitCount(w http.ResponseWriter, r *http.Request) {

	widgets, count, err := c.Store.Widget().SelectLimitCount(r.Context(), 99)
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
