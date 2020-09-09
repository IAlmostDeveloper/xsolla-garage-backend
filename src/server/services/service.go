package services

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetIdFromPath(request *http.Request) int {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	return id
}
