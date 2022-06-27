package product

import (
	"errors"
	"shop_api/pkg/config"
)

func Create(product *Product) {
	config.DB.Create(product)
}

func FindProducts(products *[]Product) {
	config.DB.Find(products)
}

func FindProduct(product *Product, conditions ...interface{}) error {
	result := config.DB.Find(product, conditions...)
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

func SaveProduct(product *Product) {
	config.DB.Save(product)
}

func Delete(product *Product) {
	config.DB.Delete(product)
}
