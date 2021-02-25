package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Connect to a database handle from a connection string
func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=nicolascamacho password=tiempo18 dbname=phrases_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
