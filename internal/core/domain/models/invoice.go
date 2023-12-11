package models

import "time"

type Invoice struct {
	ID          uint64        `gorm:"primary_key;auto_increment" json:"invoice_id"`
	CustomerID  uint64        `gorm:"not null" json:"customer_id"`
	Customer    Customer      `gorm:"foreignkey:CustomerID" json:"customer"`
	IssueDate   time.Time     `gorm:"type:date" json:"issue_date"`
	DueDate     time.Time     `gorm:"type:date" json:"due_date"`
	Subject     string        `json:"subject"`
	TotalAmount float64       `gorm:"type:decimal(12,2)" json:"total_amount"`
	Status      int           `gorm:"type:tinyint(5)" json:"status"`
	InvoiceItem []InvoiceItem `gorm:"foreignkey:InvoiceID" json:"invoice_items"`
}
