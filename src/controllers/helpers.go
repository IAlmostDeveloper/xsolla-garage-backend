package controllers

import (
	"encoding/json"
	"net/http"
)

func errorRespond(w http.ResponseWriter, r *http.Request, code int, err error) {
	respond(w, r, code, map[string]string{"error": err.Error()})
}

func respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
