package main

import (
	"github.com/julienschmidt/httprouter"
)

type protocol string

const ProtocolREST protocol = "REST"

type AppRoutes struct {
	Protocol *protocol
}

func NewRouter(p protocol) *httprouter.Router {

	switch p {
	case ProtocolREST:
		return newRestRouter()

	default:
		return newRestRouter() //should maybe panic if it can
	}

}

func newRestRouter() *httprouter.Router {

	controller := SampleController{}

	router := httprouter.New()
	router.HandlerFunc("GET", "/object", controller.List)
	router.HandlerFunc("GET", "/object/:id", controller.GetOne)
	router.HandlerFunc("POST", "/object/:id/do-something", controller.WeirdStuff)

	return router

}

func newJSONRPCRouter() *httprouter.Router {

	controller := SampleController{}

	router := httprouter.New()
	router.HandlerFunc("object.GetOne", "/jsonrpcendpoint", controller.GetOne)

	return router

}
