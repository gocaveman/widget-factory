package main

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctrl *SampleController) *httprouter.Router {

	router := httprouter.New()
	router.HandlerFunc("GET", "/object", ctrl.List)
	router.HandlerFunc("GET", "/object/:id", ctrl.GetOne)
	router.HandlerFunc("POST", "/object/:id/do-something", ctrl.WeirdStuff)

	return router

}
