package main

import (
	"github.com/brettman/smartAcGo/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/devices", handlers.GetAllDevices("this is my test"))
	}
	router.Run()
}
