package server

import (
	"fmt"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/storage/mysqlStorage"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDb(config.DbConnection)
	if err != nil {
		return err
	}
	defer db.Close()
	storage := mysqlStorage.New(db)
	server := NewServer(storage)

	port := ":8081" // ":" + os.Getenv("PORT") // for env var $PORT
	fmt.Println("Port " + port)

	return http.ListenAndServe(port, server.router)
}

func newDb(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
