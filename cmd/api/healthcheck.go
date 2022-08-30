//Filename: cmd/api/healthcheck.go

package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//create a map to hold our healthcheck data
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	//convert our map into a JSON object
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encountered a proble and could not process your request", http.StatusInternalServerError)
		return
	}
	// add a new a line to make viewing on the terminal easier
	js = append(js, '\n')
	// specify that we will server our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	//write the []byte slice containg the JSON response body
	w.Write(js)

}
