package main

import (
	"github.com/brettman/smartAcGo/internal/data"
	"github.com/brettman/smartAcGo/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	deviceService := &data.DeviceService{Str: "my string"}

	api := router.Group("/api")
	{
		api.GET("/devices", handlers.Devices(deviceService))
		api.GET("/devices/:id", handlers.Device(deviceService))
		api.GET("/devices/find/:partialId", handlers.Find(deviceService))
		api.POST("/devices", handlers.Register(deviceService))
	}
	router.Run()
}
