package handlers

import (
	"net/http"

	sac "github.com/brettman/smartacgo"
	"github.com/gin-gonic/gin"
)

// GetAllDevices - get all devices
func GetAllDevices(svc sac.DeviceService) func(c *gin.Context) {
	// todo:  should svc param be a pointer?  having some weird build issues, but they are totally due
	//  to my ignorance... need to figure out what the right approach is here

	return func(c *gin.Context) {
		svc.Devices()
		c.JSON(http.StatusOK, "got some devices")
	}
}

// GetDevice - get a single device
func GetDevice(svc *sac.DeviceService) func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "got a single device")
	}
}
