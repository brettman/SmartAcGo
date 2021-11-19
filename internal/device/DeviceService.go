package device

// DeviceService - interface to define how we deal with devices
type DeviceService interface {
	Device(serialNr string) (Device, error)
	Devices() ([]Device, error)
	Register(d Device) error
	Find(partialSerialnr string) ([]Device, error)
	AddSensorData(serialNr string, entries []SensorLogEntry) error
}
