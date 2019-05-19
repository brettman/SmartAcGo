package smartacgo

import "time"

// Device - an air conditioner which will submit data readings
type Device struct {
	SerialNr         string
	FirmwareVersion  string
	RegistrationDate time.Time
}

// DeviceService - interface to define how we deal with devices
type DeviceService interface {
	Device(serialNr string) (*Device, error)
	Devices() ([]*Device, error)
	Register(d *Device) error
	Find(partialSerialnr string) ([]*Device, error)
}
