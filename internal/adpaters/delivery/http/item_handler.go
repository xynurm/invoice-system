package http

import (
	"context"
	"invoice-system/internal/core/domain/dto"
	"invoice-system/internal/core/ports"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type itemHandlerImpl struct {
	itemUsecase ports.ItemUsecase
	validation  *validator.Validate
	wg          sync.WaitGroup
}

func NewItemHandler(itemUsecase ports.ItemUsecase) *itemHandlerImpl {
	return &itemHandlerImpl{itemUsecase, validator.New(), sync.WaitGroup{}}
}

func (h *itemHandlerImpl) CreateItemHandler(c *gin.Context) {
	var request dto.ItemRequest

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
		itemResponse, err := h.itemUsecase.CreateItemUsecase(ctx, request)
		if err != nil {
			errResponse := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}

		response := dto.SuccessResult{Code: http.StatusOK, Data: itemResponse}
		c.JSON(http.StatusOK, response)
	}()

	h.wg.Wait()
}
