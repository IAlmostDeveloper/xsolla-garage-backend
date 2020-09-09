package server

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/server/controllers"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/server/services"
)

func Start(config *services.Config) error{
	err := services.Initialize(config)
	if err != nil {
		return err
	}
	err = controllers.Handle()
	return err
}
