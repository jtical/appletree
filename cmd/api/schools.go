//Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createSchoolHandler for the "POST /v1/schools" endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new school..")
}

// createSchoolHandler for the "GET /v1/schools/:id" endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//use the "paramsfromcontext()" function to get the request context as a slice
	params := httprouter.ParamsFromContext(r.Context())
	//get the valuse of the "id" parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//Display the shcool id
	fmt.Fprintf(w, "show the details for the school %d\n", id)
}
