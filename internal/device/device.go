package device

import (
	"time"
)

// Device - an air conditioner which will submit data readings
type Device struct {
	SerialNr         string `gorm:"primaryKey;unique"`
	FirmwareVersion  string
	RegistrationDate time.Time
	SensorLogs       []SensorLogEntry `gorm:"foreignKey:SerialNr"`
}

type SensorLogEntry struct {
	SerialNr string
	Temperature  float32
	Humidity     float32
	COLevel      float32
	DeviceHealth string `gorm:size:50`
	Timestamp    time.Time
}
