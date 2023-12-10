package repository

import (
	"context"
	"invoice-system/internal/core/domain/models"
	"invoice-system/internal/core/ports"

	"gorm.io/gorm"
)

func NewCustomerRepository(db *gorm.DB) ports.CustomerRepository {
	return &repository{db}
}

func (r *repository) FindCustomersRepository(ctx context.Context) ([]models.Customer, error) {
	var err error
	var customers []models.Customer

	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return customers, tx.Error
	}

	err = tx.Find(&customers, r).Error
	if err != nil {
		tx.Rollback()
		return customers, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return customers, err
	}

	return customers, err
}

func (r *repository) GetCustomerRepository(ctx context.Context, customerID int) (models.Customer, error) {
	var err error
	var customer models.Customer

	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return customer, tx.Error
	}

	err = tx.First(&customer, "id=?", customerID).Error
	if err != nil {
		tx.Rollback()
		return customer, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return customer, err
	}

	return customer, err
}

func (r *repository) CreateCustomerRepository(ctx context.Context, customer models.Customer) (models.Customer, error) {
	var err error

	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return customer, tx.Error
	}

	err = tx.Create(&customer).Error
	if err != nil {
		tx.Rollback()
		return customer, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return customer, err
	}

	return customer, err
}
