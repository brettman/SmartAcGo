package data

import (
	"fmt"

	sac "github.com/brettman/smartacgo"
)

// DeviceService - a mongodb impl of the DeviceService interface
type DeviceService struct {
	// Db *mgo.Database
	Str string
}

// implements the DeviceService interface in root package

// Device - get a single device by serial nr
func (s *DeviceService) Device(serialNr string) (*sac.Device, error) {
	return nil, nil
}

// Devices - get all Devices
func (s *DeviceService) Devices() ([]*sac.Device, error) {
	fmt.Println("getting alll the devices from inside the repository...")
	return nil, nil
}

// Register - add a new device
func (s *DeviceService) Register(device *sac.Device) error {
	return nil
}

// Find - find all devices using a partial starting value for serialNr
func (s *DeviceService) Find(partialSerialnr string) ([]*sac.Device, error) {
	return nil, nil
}
