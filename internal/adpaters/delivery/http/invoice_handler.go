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

type invoiceHandlerImpl struct {
	invoiceUsecase ports.InvoiceUsecase
	validation     *validator.Validate
	wg             sync.WaitGroup
}

func NewInvoiceHandler(invoiceUsecase ports.InvoiceUsecase) ports.InvoiceHandler {
	return &invoiceHandlerImpl{invoiceUsecase, validator.New(), sync.WaitGroup{}}
}

func (h *invoiceHandlerImpl) FindInvoicesHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		invoicesResponse, err := h.invoiceUsecase.FindInvoicesUsecase(ctx)
		if err != nil {
			errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}
		response := dto.SuccessResult{Code: http.StatusOK, Data: invoicesResponse}
		c.JSON(http.StatusOK, response)
	}()
	h.wg.Wait()
}

func (h *invoiceHandlerImpl) CreateInvoiceHandler(c *gin.Context) {
	var request dto.InvoiceRequest

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
		invoiceResponse, err := h.invoiceUsecase.CreateInvoiceUsecase(ctx, request)
		if err != nil {
			errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}

		response := dto.SuccessResult{Code: http.StatusOK, Data: invoiceResponse}
		c.JSON(http.StatusOK, response)
	}()
	h.wg.Wait()
}

func (h *invoiceHandlerImpl) GetInvoiceHandler(c *gin.Context) {
	strID := c.Param("id")

	invoiceID, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		invoiceResponse, err := h.invoiceUsecase.GetInvoiceUsecase(ctx, invoiceID)
		if err != nil {
			errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}
		response := dto.SuccessResult{Code: http.StatusOK, Data: invoiceResponse}
		c.JSON(http.StatusOK, response)
	}()
	h.wg.Wait()
}
