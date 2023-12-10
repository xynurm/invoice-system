package models

type Customer struct {
	ID      int    `json:"customer_id" gorm:"primary_key:auto_increment"`
	Name    string `json:"name"  gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: varchar(255)"`
}
