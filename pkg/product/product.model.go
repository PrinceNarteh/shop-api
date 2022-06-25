package product

import (
	"time"
)

type Product struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Price        string    `json:"price" validate:"required"`
	SerialNumber string    `json:"serialNumber" validate:"required"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
