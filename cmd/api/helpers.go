//Filename: cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	//use the "paramsfromcontext()" function to get the request context as a slice
	params := httprouter.ParamsFromContext(r.Context())
	//get the valuse of the "id" parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	//convert our map into a JSON object
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// add a new a line to make viewing on the terminal easier
	js = append(js, '\n')
	//add the headers
	for key, value := range headers {
		w.Header()[key] = value
	}
	// specify that we will server our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//write the []byte slice containg the JSON response body
	w.Write(js)
	return nil
}
