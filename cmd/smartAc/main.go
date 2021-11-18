package main

import (
	"log"

	"github.com/brettman/smartAcGo/internal/data"
	"github.com/brettman/smartAcGo/internal/handlers"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

// App - Basic application info
type App struct {
	Name string
	Version string
}

func (app *App) Run() error{
	router := gin.Default()

	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Panic(err)
	}
	db := session.DB("smartac")

	//deviceService := &data.DeviceService{Db: db}
	deviceService:= data.NewDeviceService(db)

	api := router.Group("/api")
	{
		api.GET("/devices", handlers.Devices(deviceService))
		api.GET("/devices/:id", handlers.Device(deviceService))
		//api.GET("/devices/find/:partialId", handlers.Find(deviceService))
		api.POST("/devices", handlers.Register(deviceService))
	}
	router.Run()

	return nil
}

func main() {
	app:= &App{
		Name: "SmartAcApp",
		Version: ".0.0.1",
	}

	if err:= app.Run(); err != nil{
		log.Fatal("Error starting app")
		panic(app)
	}

}
