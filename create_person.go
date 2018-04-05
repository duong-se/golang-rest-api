package main

import (
	"encoding/json"
	"net/http"
)

// CreatePerson function to get People
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}
