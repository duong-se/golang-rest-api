package main

import (
	"encoding/json"
	"net/http"
)

// GetPeople for get all people in json
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
