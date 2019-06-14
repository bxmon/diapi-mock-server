package main

import (
	"github.com/bxmon/diapi-mock-server/controller"
	"github.com/bxmon/diapi-mock-server/router"
	"github.com/bxmon/diapi-mock-server/service"
	"github.com/bxmon/diapi-mock-server/storage"
	"github.com/gin-gonic/gin"
)

// SetUpEngine set up a server engine
func SetUpEngine(c *controller.Controller) *gin.Engine {
	engine := gin.Default()
	router.NewRouters(c, engine)
	return engine
}

func main() {
	storage := storage.NewStorage("account.db", "accountbucket")
	defer storage.BoltDB.Close()
	service := service.NewService(storage)
	controller := controller.NewController(service)
	server := SetUpEngine(controller)
	server.Run(":5555")
}
