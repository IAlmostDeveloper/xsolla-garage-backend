package services

import (
	"strconv"
	"strings"
)

func GetIdFromPath(path string) int {
	p := strings.Split(path, "/")
	res, _ := strconv.Atoi(p[2])
	return res
}
