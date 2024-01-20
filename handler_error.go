package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	// type errorResponse struct {
	// 	Error string `json:"error"`
	// }
	respondWithError(w, 500, "Internal server error")
}
