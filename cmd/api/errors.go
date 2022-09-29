package main

import (
	"fmt"
	"net/http"
)

// log error messgaes
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// we want to send JSON formatted error message
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	// create the JSON response
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// server error response
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log the eror
	app.logError(r, err)
	//prepare the message with the error
	message := "the server encountered a problem and could not process a request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// The not found response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	//create our message
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// a method not allowed response
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	//create our message
	message := fmt.Sprintf("the %s mothod is not supported for this resourc", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
