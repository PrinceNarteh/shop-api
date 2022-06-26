package order

import (
	"time"

	"shop_api/pkg/product"
	"shop_api/pkg/user"
)

type Order struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	Quantity  uint            `json:"quantity"`
	ProductId uint            `json:"productId"`
	Product   product.Product `json:"product" gorm:"foreignKey:ProductId"`
	UserId    uint            `json:"userId"`
	User      user.User       `json:"user" gorm:"foreignKey:UserId"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
