package pgsqldb

import(
	"github.com/brettman/smartAcGo/internal/device"
	"gorm.io/gorm"
	//"gorm.io/driver/postgres"
)

type deviceRepository struct {
	DB *gorm.DB
}

func NewDeviceServicePG(DB *gorm.DB) device.DeviceService {
	return &deviceRepository{DB: DB}
}

func (d deviceRepository) Device(serialNr string) (device.Device, error) {
	var item device.Device
	if result:=d.DB.First(&item, "serial_nr = ? ", serialNr); result.Error !=nil{
		return device.Device{}, result.Error
	}
	return item, nil
}

func (d deviceRepository) Devices() ([]device.Device, error) {
	var devices []device.Device
	if result:= d.DB.Find(&devices); result.Error != nil{
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
	panic("implement me")
}

func (d deviceRepository) AddSensorData(serialNr string, entries []device.SensorLogEntry) error {
	panic("implement me")
}

