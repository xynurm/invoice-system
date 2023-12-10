package models

type Item struct {
	ID   int    `json:"item_id" gorm:"primary_key:auto_increment"`
	Name string `json:"name"  gorm:"type: varchar(255)"`
	Type string `json:"type" gorm:"type: varchar(255)"`
}
