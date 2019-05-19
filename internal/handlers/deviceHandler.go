package handlers

import (
	"net/http"

	sac "github.com/brettman/smartacgo"
	"github.com/gin-gonic/gin"
)

// GetAllDevices - closure which accepts a device service and returns a gin context handler
func GetAllDevices(svc *sac.DeviceService) func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "test")
	}
}
