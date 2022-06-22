package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string `json:"name"`
	SerialNumber string `json:"serialNumber"`
}
