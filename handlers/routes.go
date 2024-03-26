package handlers

import (
	"encoding/json"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	ResponseWithJSON(w, http.StatusOK, "Hello, World!!")
}

func ResponseWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
