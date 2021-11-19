package device

// DeviceService - interface to define how we deal with devices
type DeviceService interface {
	Device(serialNr string) (Device, error)
	Devices() ([]Device, error)
	Register(d Device) error
	Find(partialSerialnr string) ([]Device, error)
}

type SensorLogEntryService interface {
	Update(serialNr string, entry SensorLogEntry) error
	UpdateBulk(serialNr string, entires []SensorLogEntry) error
}
