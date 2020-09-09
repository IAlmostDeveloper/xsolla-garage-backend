package services

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/store/mysqlStore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Initialize(config *Config) error{
	db, err := sqlx.Open("mysql", config.DbConnection)
	if err != nil {
		return err
	}
	store = mysqlStore.New(db)
	return nil
}
