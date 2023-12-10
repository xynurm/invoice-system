package ports

import (
	"context"
	"invoice-system/internal/core/domain/dto"
	"invoice-system/internal/core/domain/models"

	"github.com/gin-gonic/gin"
)

type CustomerRepository interface {
	FindCustomersRepository(ctx context.Context) ([]models.Customer, error)
	CreateCustomerRepository(ctx context.Context, customer models.Customer) (models.Customer, error)
	GetCustomerRepository(ctx context.Context, customerID int) (models.Customer, error)
}

type CustomerUsecase interface {
	CreateCustomerUsecase(ctx context.Context, customer dto.CustomerRequest) (*dto.CustomerResponse, error)
	FindCustomersUsecase(ctx context.Context) ([]models.Customer, error)
	GetCustomerUsecase(ctx context.Context, customerID int) (*dto.CustomerResponse, error)
}

type CustomerHandler interface {
	FindCustomersHandler(c *gin.Context)
	GetCustomerHandler(c *gin.Context)
	CreateCustomerHandler(c *gin.Context)
}
