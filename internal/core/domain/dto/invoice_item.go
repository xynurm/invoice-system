package dto

type InvoiceItemRequest struct {
	ItemID    uint64  `json:"item_id" validate:"required"`
	Qty       float64 `json:"qty" validate:"required,gt=0"`
	UnitPrice float64 `json:"unit_price" validate:"required,gt=0"`
}

type InvoiceItemResponse struct {
	InvoiceItemID uint64  `json:"invoice_item_id"`
	ItemID        uint64  `json:"item_id"`
	Name          string  `json:"name"`
	Qty           float64 `json:"qty"`
	UnitPrice     float64 `json:"unit_price"`
	Amount        float64 `json:"amount"`
}
