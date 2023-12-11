package repository

import (
	"context"

	"invoice-system/internal/core/domain/models"
	"invoice-system/internal/core/ports"

	"gorm.io/gorm"
)

func NewInvoiceRepository(db *gorm.DB) ports.InvoiceRepository {
	return &repository{db}

}

func (r *repository) FindInvoicesRepository(ctx context.Context) ([]models.Invoice, error) {
	var err error
	var invoices []models.Invoice

	tx := r.db.WithContext(ctx).Preload("Customer").
		Preload("InvoiceItem").
		Preload("InvoiceItem.Item").Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return invoices, tx.Error
	}

	err = tx.Find(&invoices).Error
	if err != nil {
		tx.Rollback()
		return invoices, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return invoices, err
	}

	return invoices, err
}

func (r *repository) CreateInvoiceRepository(ctx context.Context, invoice models.Invoice) (models.Invoice, error) {
	var err error

	tx := r.db.WithContext(ctx).Preload("Customer").
		Preload("InvoiceItem").
		Preload("InvoiceItem.Item").Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return invoice, tx.Error
	}

	err = tx.Create(&invoice).Error
	if err != nil {
		tx.Rollback()
		return invoice, err
	}

	// Assuming invoice.InvoiceItem is a slice of associated items
	for i := range invoice.InvoiceItem {
		// Set the InvoiceID before creating
		invoice.InvoiceItem[i].InvoiceID = invoice.ID

		// Set the ID to 0 to allow the database to auto-increment
		invoice.InvoiceItem[i].ID = 0

		err = tx.Create(&invoice.InvoiceItem[i]).Error
		if err != nil {
			tx.Rollback()
			return invoice, err
		}
	}

	err = tx.Preload("Customer").
		Preload("InvoiceItem").
		Preload("InvoiceItem.Item").
		Where("id = ?", invoice.ID).
		First(&invoice).Error

	if err != nil {
		tx.Rollback()
		return invoice, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return invoice, err
	}

	return invoice, nil
}

func (r *repository) GetInvoiceRepository(ctx context.Context, invoiceID uint64) (models.Invoice, error) {
	var err error
	var invoice models.Invoice

	tx := r.db.WithContext(ctx).Preload("Customer").
		Preload("InvoiceItem").
		Preload("InvoiceItem.Item").Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return invoice, tx.Error
	}

	err = tx.First(&invoice, "id=?", invoiceID).Error
	if err != nil {
		tx.Rollback()
		return invoice, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return invoice, err
	}

	return invoice, err
}
