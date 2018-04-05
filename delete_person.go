package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// DeletePerson function to get People
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
