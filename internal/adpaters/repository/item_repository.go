package repository

import (
	"context"
	"errors"
	"invoice-system/internal/core/domain/models"
	"invoice-system/internal/core/ports"

	"gorm.io/gorm"
)

func NewItemRepository(db *gorm.DB) ports.ItemRepository {
	return &repository{db}
}

func (r *repository) FindItemsRepository(ctx context.Context) ([]models.Item, error) {
	var err error
	var items []models.Item

	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return items, tx.Error
	}

	err = tx.Find(&items, r).Error
	if err != nil {
		tx.Rollback()
		return items, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return items, err
	}

	return items, err
}

func (r *repository) GetItemRepository(ctx context.Context, itemID int) (models.Item, error) {
	var err error
	var item models.Item

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

	err = tx.First(&item, "id=?", itemID).Error
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

func (r *repository) CreateItemRepository(ctx context.Context, item models.Item) (models.Item, error) {
	var err error

	if item.Name == "" {
		return models.Item{}, errors.New("Bad Request: Name is required")
	}

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
