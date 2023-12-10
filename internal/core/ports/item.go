package ports

import (
	"context"
	"invoice-system/internal/core/domain/dto"
	"invoice-system/internal/core/domain/models"
)

type ItemRepository interface {
	FindItemsRepository(ctx context.Context) ([]models.Item, error)
	CreateItemRepository(ctx context.Context, item models.Item) (models.Item, error)
}

type ItemUsecase interface {
	CreateItemUsecase(ctx context.Context, item dto.ItemRequest) (*dto.ItemResponse, error)
	FindItemsUsecase(ctx context.Context) ([]models.Item, error)
}
