package handlers

import (
	"net/http"

	sac "github.com/brettman/smartacgo"
	"github.com/gin-gonic/gin"
)

// Devices - get all devices
func Devices(svc sac.DeviceService) func(c *gin.Context) {
	// todo:  should svc param be a pointer?  having some weird build issues, but they are totally due
	//  to my ignorance... need to figure out what the right approach is here

	return func(c *gin.Context) {
		svc.Devices()
		c.JSON(http.StatusOK, "got some devices")
	}
}

// Device - get a single device
func Device(svc sac.DeviceService) func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "got a single device")
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
		c.JSON(http.StatusOK, "posting a new device")
	}
}
