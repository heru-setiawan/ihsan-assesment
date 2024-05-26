package datastore

import "gorm.io/gorm"

func NewDatabase(db *gorm.DB) *database {
	return &database{
		db: db,
	}
}

type database struct {
	db *gorm.DB
}
