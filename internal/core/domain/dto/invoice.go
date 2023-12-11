package dto

type InvoiceRequest struct {
	CustomerID  uint64               `json:"customer_id" validate:"required"`
	IssueDate   string               `json:"issue_date" validate:"required"`
	DueDate     string               `json:"due_date" validate:"required"`
	Subject     string               `json:"subject" validate:"required"`
	TotalAmount float64              `json:"total_amount" validate:"required,gt=0"`
	Status      int                  `json:"status" validate:"required"`
	Items       []InvoiceItemRequest `json:"items" validate:"required,dive,required"`
}

type InvoiceResponse struct {
	ID           uint64                `json:"id"`
	CustomerID   uint64                `json:"customer_id"`
	CustomerName string                `json:"customer_name"`
	IssueDate    string                `json:"issue_date"`
	DueDate      string                `json:"due_date"`
	Subject      string                `json:"subject"`
	TotalAmount  float64               `json:"total_amount"`
	Status       int                   `json:"status"`
	Items        []InvoiceItemResponse `json:"items"`
}
