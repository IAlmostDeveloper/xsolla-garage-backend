package services

import (
	"strconv"
	"strings"
	"github.com/gorilla/mux"
)

func GetIdFromPath(path string) int {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	return id
}
