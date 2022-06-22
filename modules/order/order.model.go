package order

import (
	"gorm.io/gorm"

	"shop_api/modules/product"
	"shop_api/modules/user"
)

type Order struct {
	gorm.Model
	ProductId uint            `json:"productId"`
	Product   product.Product `gorm:"foreignKey:ProductId"`
	UserId    uint            `json:"userId"`
	User      user.User       `gorm:"foreignKey:UserId"`
}
