//Filename: cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Define a new type name envelope
type envelope map[string]interface{}

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

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//convert our map into a JSON object
	js, err := json.MarshalIndent(data, "", "\t")
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

// dst stores decoded struct
func (app *application) readJSON(W http.ResponseWriter, r *http.Request, dst interface{}) error {
	//decode the request body into the target destination dst
	err := json.NewDecoder(r.Body).Decode(dst)
	//check for a bad request
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		//switch to check for the errors
		switch {
		//check for syntax errors
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON(at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		//Check for the wrong type pass by the user
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		//Emtpy body
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		//Pass non-nil pointer error
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
			//default
		default:
			return err
		}

	}
	return nil
}
