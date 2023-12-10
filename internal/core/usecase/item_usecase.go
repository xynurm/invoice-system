package usecase

import (
	"context"
	"invoice-system/internal/core/domain/dto"
	"invoice-system/internal/core/domain/models"
	"invoice-system/internal/core/ports"
)

type itemUsecaseImpl struct {
	itemRepository ports.ItemRepository
}

func NewItemUsecase(itemRepository ports.ItemRepository) ports.ItemUsecase {
	return &itemUsecaseImpl{itemRepository}
}

func (u *itemUsecaseImpl) CreateItemUsecase(ctx context.Context, req dto.ItemRequest) (*dto.ItemResponse, error) {
	createItem := models.Item{
		Name: req.Name,
		Type: req.Type,
	}

	data, err := u.itemRepository.CreateItemRepository(ctx, createItem)

	if err != nil {
		return nil, err
	}

	response := &dto.ItemResponse{
		ID:   data.ID,
		Name: data.Name,
		Type: data.Type,
	}

	return response, err
}

func (u *itemUsecaseImpl) FindItemsUsecase(ctx context.Context) ([]models.Item, error) {
	data, err := u.itemRepository.FindItemsRepository(ctx)

	if err != nil {
		return nil, err
	}

	return data, err
}
