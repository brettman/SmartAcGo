package handlers

import (
	"log"
	"net/http"

	sac "github.com/brettman/smartacgo"
	"github.com/gin-gonic/gin"
)

// Devices - get all devices
func Devices(svc sac.DeviceService) func(c *gin.Context) {
	return func(c *gin.Context) {
		devices, err := svc.Devices()
		if err != nil {
			log.Panic(err)
		}

		c.JSON(http.StatusOK, devices)
	}
}

// Device - get a single device
func Device(svc sac.DeviceService) func(c *gin.Context) {

	return func(c *gin.Context) {
		id := c.Param("id")
		device, err := svc.Device(id)
		if err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, device)
	}
}

// Find - pass a starting substring of a serial nr and find a device
func Find(svc sac.DeviceService) func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "searching for devices with a start string serial nr")
	}
}

// Register - add a new device to the db
func Register(svc sac.DeviceService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var device sac.Device
		err := c.ShouldBindJSON(&device)
		if err != nil {
			log.Panic(err)
		}
		err = svc.Register(device)
		if err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, device)
	}
}
