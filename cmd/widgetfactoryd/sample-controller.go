package main

import (
	"fmt"
	"net/http"
)

type SampleController struct {
	Message string
}

func NewSampleController() *SampleController {

	return &SampleController{Message: "Hello World"}

}

//e.g. a list of the resource
func (c *SampleController) List(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, c.Message)
}

//get one resource
func (c *SampleController) GetOne(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	_, _ = fmt.Fprintf(w, c.Message+" id is %v", id)
}

func (c *SampleController) WeirdStuff(w http.ResponseWriter, r *http.Request) {

	//do whatever weird shit we need to do
	_, _ = fmt.Fprintf(w, c.Message)
}
