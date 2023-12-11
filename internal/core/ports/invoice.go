package ports

import (
	"context"
	"invoice-system/internal/core/domain/dto"
	"invoice-system/internal/core/domain/models"

	"github.com/gin-gonic/gin"
)

type InvoiceRepository interface {
	CreateInvoiceRepository(ctx context.Context, invoice models.Invoice) (models.Invoice, error)
	FindInvoicesRepository(ctx context.Context) ([]models.Invoice, error)
	GetInvoiceRepository(ctx context.Context, invoiceID uint64) (models.Invoice, error)
}

type InvoiceUsecase interface {
	CreateInvoiceUsecase(ctx context.Context, req dto.InvoiceRequest) (*dto.InvoiceResponse, error)
	FindInvoicesUsecase(ctx context.Context) ([]models.Invoice, error)
	GetInvoiceUsecase(ctx context.Context, invoiceID uint64) (*dto.InvoiceResponse, error)
}

type InvoiceHandler interface {
	CreateInvoiceHandler(c *gin.Context)
	FindInvoicesHandler(c *gin.Context)
	GetInvoiceHandler(c *gin.Context)
}
