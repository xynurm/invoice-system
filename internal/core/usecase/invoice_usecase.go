package usecase

import (
	"context"
	"invoice-system/internal/core/domain/dto"
	"invoice-system/internal/core/domain/models"
	"invoice-system/internal/core/ports"
	"time"
)

type invoiceUsecaseImpl struct {
	invoiceRepository ports.InvoiceRepository
}

func NewInvoiceUsecase(invoiceRepository ports.InvoiceRepository) ports.InvoiceUsecase {
	return &invoiceUsecaseImpl{invoiceRepository}
}

func parseDate(dateString string) (time.Time, error) {
	// Define the expected date layout
	layout := "02/01/2006" // DD/MM/YYYY

	// Parse the date string
	return time.Parse(layout, dateString)
}

func (u *invoiceUsecaseImpl) CreateInvoiceUsecase(ctx context.Context, req dto.InvoiceRequest) (*dto.InvoiceResponse, error) {
	// Parse the issue and due dates
	issueDate, err := parseDate(req.IssueDate)
	if err != nil {
		return nil, err
	}

	dueDate, err := parseDate(req.DueDate)
	if err != nil {
		return nil, err
	}

	// Create an Invoice object from the request DTO
	createInvoice := models.Invoice{
		CustomerID:  req.CustomerID,
		IssueDate:   issueDate,
		DueDate:     dueDate,
		Subject:     req.Subject,
		TotalAmount: req.TotalAmount,
		Status:      req.Status,
	}

	// Create InvoiceItems from the request DTO
	var invoiceItems []models.InvoiceItem
	for _, item := range req.Items {
		invoiceItem := models.InvoiceItem{
			ItemID:    item.ItemID,
			Qty:       item.Qty,
			UnitPrice: item.UnitPrice,
			Amount:    item.Qty * item.UnitPrice,
		}
		invoiceItems = append(invoiceItems, invoiceItem)
	}

	// Associate InvoiceItems with the Invoice
	createInvoice.InvoiceItem = invoiceItems

	// Call the repository to create the invoice and its items
	data, err := u.invoiceRepository.CreateInvoiceRepository(ctx, createInvoice)
	if err != nil {
		return nil, err
	}

	issueDateStr := data.IssueDate.Format("02/01/2006")
	dueDateStr := data.DueDate.Format("02/01/2006")

	// Map the created data to response DTO
	response := &dto.InvoiceResponse{
		ID:           data.ID,
		CustomerID:   data.CustomerID,
		CustomerName: data.Customer.Name,
		IssueDate:    issueDateStr,
		DueDate:      dueDateStr,
		Subject:      data.Subject,
		TotalAmount:  data.TotalAmount,
		Status:       data.Status,
		Items:        mapInvoiceItemsToDTO(data.InvoiceItem),
	}

	return response, nil
}
func mapInvoiceItemsToDTO(items []models.InvoiceItem) []dto.InvoiceItemResponse {
	var responseItems []dto.InvoiceItemResponse
	for _, item := range items {
		responseItem := dto.InvoiceItemResponse{
			InvoiceItemID: item.ID,
			ItemID:        item.ItemID,
			Name:          item.Item.Name,
			Qty:           item.Qty,
			UnitPrice:     item.UnitPrice,
			Amount:        item.Amount,
		}
		responseItems = append(responseItems, responseItem)
	}
	return responseItems
}

func (u *invoiceUsecaseImpl) FindInvoicesUsecase(ctx context.Context) ([]models.Invoice, error) {
	invoices, err := u.invoiceRepository.FindInvoicesRepository(ctx)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (u *invoiceUsecaseImpl) GetInvoiceUsecase(ctx context.Context, invoiceID uint64) (*dto.InvoiceResponse, error) {
	invoice, err := u.invoiceRepository.GetInvoiceRepository(ctx, invoiceID)
	if err != nil {
		return nil, err
	}

	issueDateStr := invoice.IssueDate.Format("02/01/2006")
	dueDateStr := invoice.DueDate.Format("02/01/2006")

	response := &dto.InvoiceResponse{
		ID:          invoice.ID,
		CustomerID:  invoice.CustomerID,
		IssueDate:   issueDateStr,
		DueDate:     dueDateStr,
		Subject:     invoice.Subject,
		TotalAmount: invoice.TotalAmount,
		Status:      invoice.Status,
		Items:       mapInvoiceItemsToDTO(invoice.InvoiceItem),
	}

	return response, nil
}
