package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConnOptions struct {
	DBName string
}

func (o *DBConnOptions) SQLiteDSN() string {
	if o.DBName == "" {
		return ":memory:"
	}
	return o.DBName
}

func NewSQLiteGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(opt.SQLiteDSN()), &gorm.Config{})
}
