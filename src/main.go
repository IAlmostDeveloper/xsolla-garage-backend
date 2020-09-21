package main

import (
	"database/sql"
	"flag"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/server"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
	"log"
)

func main() {
	flag.Parse()
	config := &server.Config{}

	viper.SetConfigName("configs")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
	}
	if err := migrate(config.DbConnection); err != nil {
		log.Fatalf("migration error: %s", err)
	}
	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}

func migrate(dbConnection string) error {
	command := "up"
	dir := "./migrations"
	if db, err := sql.Open("mysql", dbConnection); err != nil {
		return err
	} else{
		defer db.Close()
		if err := goose.SetDialect("mysql"); err != nil {
			return err
		}
		if err := goose.Run(command, db, dir); err != nil {
			return err
		}
		return nil
	}
}
