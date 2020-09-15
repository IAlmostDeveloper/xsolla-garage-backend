package controllers

import (
	"encoding/json"
	"net/http"
)

func errorJsonRespond(w http.ResponseWriter, code int, err error) {
	respondJson(w, code, map[string]string{"error": err.Error()})
}

func respondJson(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
