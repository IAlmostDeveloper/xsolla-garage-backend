package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/store/mysqlStore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func Initialize(config *Config) error{
	db, err := sqlx.Open("mysql", config.DbConnection)
	if err != nil {
		return err
	}
	store = mysqlStore.New(db)
	return nil
}

func GetIdFromPath(request *http.Request) int {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	return id
}
