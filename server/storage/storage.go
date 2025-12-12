package storage

import (
	"fmt"
	"strings"

	"github.com/qwganker/boring/conf"

	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var (
	dbInstance *gorm.DB
)

func GetDBInstance() *gorm.DB {
	return dbInstance
}

func Init(cfg conf.DBConfig) error {

	var err error

	switch strings.ToLower(cfg.Type) {
	case "", "sqlite":
		dbInstance, err = InitSQLite(cfg.DSN)
	default:
		return fmt.Errorf("storage: unsupported db type %q", cfg.Type)
	}

	if err != nil {
		return fmt.Errorf("storage: failed to initialize db, %w", err)
	}

	return nil
}
