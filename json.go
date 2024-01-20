package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		fmt.Println("Responding with 5xx error: ", msg)
		return
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}
