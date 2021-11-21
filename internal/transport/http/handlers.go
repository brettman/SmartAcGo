package http

import (
	"fmt"
	"github.com/brettman/smartAcGo/internal/device"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// DeviceHandler - http handler methods for device operations
type DeviceHandler struct {
	DeviceService device.DeviceService
}

// NewDeviceHandler - init new DeviceHnadler
func NewDeviceHandler(deviceService device.DeviceService) *DeviceHandler {
	return &DeviceHandler{DeviceService: deviceService}
}

// SetupRoutes - Setup device routes
func (handler *DeviceHandler) SetupRoutes() error {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/devices", handler.Register)
		api.POST("/devices/:serialNr/sensordata", handler.SensorData)
	}
	adminApi := router.Group("/api/admin")
	{
		adminApi.GET("/devices", handler.Devices)
		adminApi.GET("/devices/:serialNr", handler.Device)
	}

	err := router.Run()
	if err != nil {
		log.Panic(err)
	}
	return nil
}

// Devices - get all devices
func (h *DeviceHandler) Devices(c *gin.Context) {
	log.Println("we're in the new method")
	devices, err := h.DeviceService.Devices()
	if err != nil {
		log.Panic(err)
	}

	c.JSON(http.StatusOK, devices)
}

// Device - get a device by serialNr
func (h *DeviceHandler) Device(c *gin.Context) {
	id := c.Param("serialNr")
	fmt.Println(id)
	acdevice, err := h.DeviceService.Device(id)
	if err != nil {
		log.Panic(err)
	}
	c.JSON(http.StatusOK, acdevice)
}

// Register - add a new device to the db
func (h *DeviceHandler) Register(c *gin.Context) {

	var acDevice device.Device

	// todo: handle deserialization errors
	err := c.ShouldBindJSON(&acDevice)
	acDevice.RegistrationDate = time.Now()
	if err != nil {
		log.Panic(err)
	}
	_, err = h.DeviceService.Register(acDevice)
	if err != nil {
		log.Panic(err)
	}
	c.JSON(http.StatusOK, acDevice)
}

func (h *DeviceHandler) SensorData(c *gin.Context) {
	var id = c.Param("serialNr")
	var sensorLogEntry []device.SensorLogEntry

	err:= c.ShouldBindJSON(&sensorLogEntry); if err !=nil{
		log.Println("invalid sensor data received")
		log.Println(err)
		err := h.DeviceService.AddSensorData(id, sensorLogEntry);if err!=nil{
			log.Panic(err)
		}
	}
}
