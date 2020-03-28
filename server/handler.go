package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-elastic-hotels/services"
	"net/http"
)

const (
	ConnectionError  = "Could not establish connection with Elastic Server"
	searchName       = "name"
	ESSearchError    = "Elastic server could not respond"
	jsonMarshalError = "Could not marshal json"
	writeError       = "could not write response"
)

func GetHomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "Welcome to Hotel Api")
	}
}

func SearchHotelHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		esClient, err := services.ConnectToESServer()
		if err != nil {
			http.Error(w, ConnectionError, http.StatusInternalServerError)
			return
		}

		params := mux.Vars(r)["name"]

		hotels, err := services.SearchForResult(ctx, esClient, searchName, params)
		if err != nil {
			http.Error(w, ESSearchError, http.StatusInternalServerError)
			return
		}

		hotelsByte, err := json.MarshalIndent(&hotels, "", " ")
		if err != nil {
			http.Error(w, jsonMarshalError, http.StatusInternalServerError)
			return
		}

		_, err = w.Write(hotelsByte)
		if err != nil {
			http.Error(w, writeError, http.StatusInternalServerError)
			return
		}
	}
}
