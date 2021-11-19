package device

import "time"

// Device - an air conditioner which will submit data readings
type Device struct {
	SerialNr         string `bson:"serialnr"`
	FirmwareVersion  string
	RegistrationDate time.Time
	SensorLogs       []SensorLogEntry
}

type SensorLogEntry struct {
	SerialNr     string
	Temperature  float32
	Humidity     float32
	COLevel      float32
	DeviceHealth string
	Timestamp    time.Time
}
