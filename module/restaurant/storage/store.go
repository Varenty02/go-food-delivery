package restaurantstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

//param:db output:sqlStore
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
