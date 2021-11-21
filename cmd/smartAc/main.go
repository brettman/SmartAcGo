package main

import (
	"github.com/brettman/smartAcGo/internal/data/pgsqldb"
	"github.com/brettman/smartAcGo/internal/transport/http"
	"log"

	"gopkg.in/mgo.v2"
)

// App - Basic application info
type App struct {
	Name string
	Version string
}

func OpenMongoDb() (*mgo.Database, error){
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Panic(err)
	}
	db := session.DB("smartac")

	return db, nil
}

func (app *App) Run() error{

	//db, _ := OpenMongoDb()
	//devService:= mongodb.NewDeviceServiceMongo(db)

	db , err:= pgsqldb.NewDatabase(); if err != nil{
		panic(err)
	}

	devService:= pgsqldb.NewDeviceServicePG(db)
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
