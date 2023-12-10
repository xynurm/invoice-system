package usecase

import (
	"context"
	"invoice-system/internal/core/domain/dto"
	"invoice-system/internal/core/domain/models"
	"invoice-system/internal/core/ports"
)

type customerUsecaseImpl struct {
	customerRepository ports.CustomerRepository
}

func NewCustomerUsecase(customerRepository ports.CustomerRepository) ports.CustomerUsecase {
	return &customerUsecaseImpl{customerRepository}
}

func (u *customerUsecaseImpl) CreateCustomerUsecase(ctx context.Context, req dto.CustomerRequest) (*dto.CustomerResponse, error) {
	createCustomer := models.Customer{
		Name:    req.Name,
		Address: req.Address,
	}

	data, err := u.customerRepository.CreateCustomerRepository(ctx, createCustomer)

	if err != nil {
		return nil, err
	}

	response := &dto.CustomerResponse{
		ID:      data.ID,
		Name:    data.Name,
		Address: data.Address,
	}

	return response, err
}

func (u *customerUsecaseImpl) FindCustomersUsecase(ctx context.Context) ([]models.Customer, error) {
	data, err := u.customerRepository.FindCustomersRepository(ctx)

	if err != nil {
		return nil, err
	}

	return data, err
}

func (u *customerUsecaseImpl) GetCustomerUsecase(ctx context.Context, customerID int) (*dto.CustomerResponse, error) {
	data, err := u.customerRepository.GetCustomerRepository(ctx, customerID)

	if err != nil {
		return nil, err
	}

	response := &dto.CustomerResponse{
		ID:      data.ID,
		Name:    data.Name,
		Address: data.Address,
	}

	return response, err
}
