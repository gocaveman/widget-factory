package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gocaveman/widget-factory/cmd/widgetfactoryd/models"
)

type WidgetController struct {
	DB *sql.DB
}

func NewWidgetController(DB *sql.DB) *WidgetController {

	return &WidgetController{
		DB: DB,
	}

}

func (c *WidgetController) List(w http.ResponseWriter, r *http.Request) {

	widgets, err := models.Widgets().All(r.Context(), c.DB)
	if err != nil {
		writeError(w, err.Error())
	}

	writeJSON(w, widgets)

}

func writeError(w http.ResponseWriter, errorString string) {

	w.WriteHeader(500)
	_, _ = w.Write([]byte(errorString))

}

func writeJSON(w http.ResponseWriter, jsonData interface{}) {

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(jsonData)

}
