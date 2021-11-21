package pgsqldb

import (
	"fmt"
	"github.com/brettman/smartAcGo/internal/device"
	"gorm.io/gorm"
)

type deviceRepository struct {
	DB *gorm.DB
}

func NewDeviceServicePG(DB *gorm.DB) device.DeviceService {
	return &deviceRepository{DB: DB}
}

func (d deviceRepository) Device(serialNr string) (device.Device, error) {
	var item device.Device
	if result:=d.DB.Preload("SensorLogs").Find(&item, "serial_nr = ? ", serialNr); result.Error !=nil{
		return device.Device{}, result.Error
	}
	return item, nil
}

func (d deviceRepository) Devices() ([]device.Device, error) {
	var devices []device.Device
	if result:= d.DB.Preload("sensor_log_entries").Find(&devices); result.Error != nil{
		return []device.Device{}, result.Error
	}
	return devices, nil
}

func (d deviceRepository) Register(item device.Device) (device.Device, error) {
	if result := d.DB.Save(item); result.Error != nil{
		return device.Device{}, result.Error
	}
	return item, nil
}

func (d deviceRepository) Find(partialSerialnr string) ([]device.Device, error) {
	var items []device.Device
	var param = fmt.Sprintf("%%%s%%", partialSerialnr)
	fmt.Printf("=-=-=-=-=-=- param is %s", param)
	if result:=d.DB.Where("serial_nr LIKE ?", param).Find(&items); result.Error !=nil{
		return []device.Device{}, result.Error
	}
	return items, nil
}

func (d deviceRepository) AddSensorData(serialNr string, entries []device.SensorLogEntry) error {
	var item device.Device
	if result:=d.DB.Find(&item, "serial_nr = ? ", serialNr); result.Error !=nil{
		return result.Error
	}
	if item.SensorLogs == nil{
		fmt.Printf("it is nil, initializing")
		item.SensorLogs = []device.SensorLogEntry{}
	}
	fmt.Println("submitted entries look like:")
	fmt.Println(entries)
	fmt.Printf("Sensorlogs look like:")
	fmt.Println(item.SensorLogs)
	fmt.Println("appending entries to the sensorlogs")

	result := append(item.SensorLogs, entries...)

	item.SensorLogs = result
	fmt.Println(item.SensorLogs)

	if result := d.DB.Save(&item); result.Error != nil {
		panic(result.Error)
	}
	return nil
}

