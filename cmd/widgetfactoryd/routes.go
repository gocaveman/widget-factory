package main

import (
	"github.com/julienschmidt/httprouter"
)

func NewSampleRouter(ctrl *SampleController) *httprouter.Router {

	router := httprouter.New()
	router.HandlerFunc("GET", "/", ctrl.List)
	router.HandlerFunc("GET", "/object", ctrl.List)
	router.HandlerFunc("GET", "/object/:id", ctrl.GetOne)
	router.HandlerFunc("POST", "/object/:id/do-something", ctrl.WeirdStuff)

	return router

}

func NewWidgetRouter(ctrl *WidgetController) *httprouter.Router {

	r := httprouter.New()
	r.HandlerFunc("GET", "/widget", ctrl.List)

	return r

}
