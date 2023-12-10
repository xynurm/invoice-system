package http

import (
	"context"
	"invoice-system/internal/core/domain/dto"
	"invoice-system/internal/core/ports"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type customerHandlerImpl struct {
	customerUsecase ports.CustomerUsecase
	validation      *validator.Validate
	wg              sync.WaitGroup
}

func NewCustomerHandler(customerUsecase ports.CustomerUsecase) ports.CustomerHandler {
	return &customerHandlerImpl{customerUsecase, validator.New(), sync.WaitGroup{}}
}

func (h *customerHandlerImpl) FindCustomersHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		customersResponse, err := h.customerUsecase.FindCustomersUsecase(ctx)
		if err != nil {
			errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}
		response := dto.SuccessResult{Code: http.StatusOK, Data: customersResponse}
		c.JSON(http.StatusOK, response)
	}()
	h.wg.Wait()

}

func (h *customerHandlerImpl) CreateCustomerHandler(c *gin.Context) {
	var request dto.CustomerRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	// Validate request input using go-playground/validator.
	if err := h.validation.Struct(request); err != nil {
		errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		customerResponse, err := h.customerUsecase.CreateCustomerUsecase(ctx, request)
		if err != nil {
			errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}

		response := dto.SuccessResult{Code: http.StatusOK, Data: customerResponse}
		c.JSON(http.StatusOK, response)
	}()
	h.wg.Wait()
}

func (h *customerHandlerImpl) GetCustomerHandler(c *gin.Context) {
	strID := c.Param("id")

	customerID, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		customersResponse, err := h.customerUsecase.GetCustomerUsecase(ctx, customerID)
		if err != nil {
			errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}
		response := dto.SuccessResult{Code: http.StatusOK, Data: customersResponse}
		c.JSON(http.StatusOK, response)
	}()
	h.wg.Wait()

}
