package main

import (
	"fmt"
	"os"

	"github.com/bxmon/diapi-mock-server/controller"
	"github.com/bxmon/diapi-mock-server/router"
	"github.com/bxmon/diapi-mock-server/service"
	"github.com/bxmon/diapi-mock-server/storage"
	"github.com/gin-gonic/gin"
)

// SetUpEngine set up a server engine
func SetUpEngine(c *controller.Controller) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(gin.Logger())
	router.NewRouters(c, engine)
	return engine
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":5000"
	}

	storage := storage.NewStorage("account.db", "accountbucket")
	defer storage.BoltDB.Close()
	service := service.NewService(storage)
	controller := controller.NewController(service)
	server := SetUpEngine(controller)
	fmt.Printf("Start gin server. Listen on port%s\n", ":5000")
	server.Run(":5000")
}
