package device

import (
	"gorm.io/gorm"
	"time"
)

// Device - an air conditioner which will submit data readings
type Device struct {
	gorm.Model
	SerialNr         string `gorm:primaryKey`
	FirmwareVersion  string
	RegistrationDate time.Time
	SensorLogs       []SensorLogEntry
}

type SensorLogEntry struct {
	gorm.Model
	SerialNr     string
	Temperature  float32
	Humidity     float32
	COLevel      float32
	DeviceHealth string `gorm:size:50`
	Timestamp    time.Time
}
