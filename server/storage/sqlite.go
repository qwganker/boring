package storage

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        dsn,
	}, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("storage: open gorm sqlite: %w", err)
	}
	return db, nil
}
