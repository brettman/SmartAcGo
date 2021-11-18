package data

import (
	"log"

	sac "github.com/brettman/smartacgo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DeviceService - a mongodb impl of the DeviceService interface
type DeviceService struct {
	Db *mgo.Database
}

func NewDeviceService(db *mgo.Database) *DeviceService {
	return &DeviceService{
		db,
		}
}

// Device - get a single device by serial nr
func (s *DeviceService) Device(serialNr string) (sac.Device, error) {
	return sac.Device{}, nil
}

// Devices - get all Devices
func (s *DeviceService) Devices() ([]sac.Device, error) {
	devices := []sac.Device{}
	err := s.Db.C("devices").Find(bson.M{}).All(&devices)
	if err != nil {
		panic(err)
	}
	return devices, nil
}

// Register - add a new device
func (s *DeviceService) Register(device sac.Device) error {
	err := s.Db.C("devices").Insert(&device)
	if err != nil {
		log.Panic(err)
	}
	return err
}

// Find - find all devices using a partial starting value for serialNr
func (s *DeviceService) Find(partialSerialnr string) ([]sac.Device, error) {
	return nil, nil
}
