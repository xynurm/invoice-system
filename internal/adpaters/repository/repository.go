package repository

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new instance of the repository with the given database connection.
func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// DB returns the underlying database connection.
func (r *repository) DB() *gorm.DB {
	return r.db
}

// SetDB sets a new database connection.
func (r *repository) SetDB(db *gorm.DB) {
	r.db = db
}
