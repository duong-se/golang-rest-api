package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPerson function to get People
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}
