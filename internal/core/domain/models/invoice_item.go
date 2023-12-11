package models

type InvoiceItem struct {
	ID        uint64  `gorm:"primary_key;auto_increment" json:"invoice_item_id"`
	InvoiceID uint64  `json:"invoice_id"`
	Invoice   Invoice `gorm:"foreignkey:InvoiceID" json:"-"`
	ItemID    uint64  `json:"item_id"`
	Qty       float64 `gorm:"type:decimal(12,2)" json:"qty"`
	UnitPrice float64 `gorm:"type:decimal(12,2)" json:"unit_price"`
	Item      Item    `gorm:"foreignkey:ItemID" json:"item"`
	Amount    float64 `gorm:"type:decimal(12,2)" json:"amount"`
}
