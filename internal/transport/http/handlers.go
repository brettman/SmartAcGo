package http

import (
	"github.com/brettman/smartAcGo/internal/device"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// DeviceHandler - http handler methods for device operations
type DeviceHandler struct{

	DeviceService device.DeviceService
}

// NewDeviceHandler - init new DeviceHnadler
func NewDeviceHandler(deviceService device.DeviceService) *DeviceHandler {
	return &DeviceHandler{DeviceService: deviceService}
}

func (handler *DeviceHandler) SetupRoutes() error{
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/devices", handler.Devices)
		api.GET("/devices/:serialNr", handler.Device)
		api.POST("/devices", handler.Register)
	}

	router.Run()
	return nil
}

func(h *DeviceHandler) Devices(c *gin.Context) {
	log.Println("we're in the new method")
	devices, err := h.DeviceService.Devices()
	if err != nil {
		log.Panic(err)
	}

	c.JSON(http.StatusOK, devices)
}

func(h *DeviceHandler) Device(c *gin.Context){
	id := c.Param("serialNr")
	device, err := h.DeviceService.Device(id); if err !=nil {
		log.Panic(err)
	}
	c.JSON(http.StatusOK, device)
}

// Register - add a new device to the db
func(h *DeviceHandler) Register(c *gin.Context){

	var device device.Device
	err := c.ShouldBindJSON(&device)
	device.RegistrationDate = time.Now()
	if err != nil {
		log.Panic(err)
	}
	err = h.DeviceService.Register(device)
	if err != nil {
		log.Panic(err)
	}
	c.JSON(http.StatusOK, device)
}
