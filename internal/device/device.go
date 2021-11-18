package device

import "time"

// Device - an air conditioner which will submit data readings
type Device struct {
	SerialNr         string `bson:"serialnr"`
	FirmwareVersion  string
	RegistrationDate time.Time
}
