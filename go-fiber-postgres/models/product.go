package models

import "gorm.io/gorm"

type Product struct{
	SKU          string   `json:"sku"`
	Title        *string   `json:"title"`
	Price        float64  `json:"price"`
	UpdateStatus *bool     `json:"updateStatus"`
	BlueFlag     bool     `json:"blueFlag"`
	IsVat        bool     `json:"isVat"`
	IsFixPrice   bool     `json:"isFixPrice"`
	Promotion    []string `json:"promotion"`
}

func MigrateProduct(db *gorm.DB) error {
	err := db.AutoMigrate(&Product{})
	return err 
}