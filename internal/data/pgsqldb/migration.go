package pgsqldb

import (
	"github.com/brettman/smartAcGo/internal/device"
	"gorm.io/gorm"
)

// MigratedDB - creates comment table with a db migration
func MigrateDB(db *gorm.DB) {
	 db.AutoMigrate(&device.Device{})
	 db.AutoMigrate(&device.SensorLogEntry{})
}
