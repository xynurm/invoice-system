package repository

import (
	"context"
	"invoice-system/internal/core/domain/models"
	"invoice-system/internal/core/ports"

	"gorm.io/gorm"
)

func NewItemRepository(db *gorm.DB) ports.ItemRepository {
	return &repository{db}
}

func (r *repository) CreateItemRepository(ctx context.Context, item models.Item) (models.Item, error) {
	var err error

	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return item, tx.Error
	}

	err = tx.Create(&item).Error
	if err != nil {
		tx.Rollback()
		return item, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return item, err
	}

	return item, err
}
