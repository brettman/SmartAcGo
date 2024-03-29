package mongodb

import (
	"fmt"
	"github.com/brettman/smartAcGo/internal/device"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DeviceService - a mongodb impl of the DeviceService interface
type deviceRepository struct {
	db *mgo.Database
}


// NewDeviceServiceMongo - returns an mongodb specific implementation of the device.DeviceService interface
func NewDeviceServiceMongo(db *mgo.Database) device.DeviceService{
	return &deviceRepository{
		db,
		}
}

// AddSensorData - update a device's sensordata log
func (s *deviceRepository) AddSensorData(serialNr string, entries []device.SensorLogEntry) error {
	var acDevice device.Device
	err := s.db.C("devices").Find(bson.M{"serialnr": serialNr}).One(&acDevice); if err != nil {
		fmt.Printf("Serialnr %s: %s\n", serialNr, err)
		return err
	}
	if acDevice.SensorLogs == nil {
		acDevice.SensorLogs = []device.SensorLogEntry{}
	}
	acDevice.SensorLogs = append(acDevice.SensorLogs, entries...)
	changeInfo, err :=s.db.C("devices").UpsertId(bson.M{"serialnr": acDevice.SerialNr}, acDevice); if err !=nil {
		log.Panic(err)
	}
	log.Println(changeInfo)
	return nil
}

// Device - get a single device by serial nr
func (s *deviceRepository) Device(serialNr string) (device.Device, error) {
	var device device.Device
	err := s.db.C("devices").Find(bson.M{"serialnr": serialNr}).One(&device); if err != nil {
		fmt.Printf("Serialnr %s: %s\n", serialNr, err)
	}
	return device, nil
}

// Devices - get all Devices
func (s *deviceRepository) Devices() ([]device.Device, error) {
	var devices []device.Device
	err := s.db.C("devices").Find(bson.M{}).All(&devices)
	if err != nil {
		panic(err)
	}
	return devices, nil
}

// Register - add a new device
func (s *deviceRepository) Register(device device.Device) (device.Device, error) {
	err := s.db.C("devices").Insert(&device)
	if err != nil {
		log.Panic(err)
	}
	return device, err
}

// Find - find all devices using a partial starting value for serialNr
func (s *deviceRepository) Find(partialSerialnr string) ([]device.Device, error) {
	return nil, nil
}
