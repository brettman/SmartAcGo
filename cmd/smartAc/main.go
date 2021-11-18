package main

import (
	"github.com/brettman/smartAcGo/internal/transport/http"
	"log"

	"github.com/brettman/smartAcGo/internal/data"
	"gopkg.in/mgo.v2"
)

// App - Basic application info
type App struct {
	Name string
	Version string
}

func (app *App) Run() error{

	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Panic(err)
	}
	db := session.DB("smartac")

	devService:= data.NewDeviceServiceMongo(db)
	handler := http.NewDeviceHandler(devService)
	handler.SetupRoutes()

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
