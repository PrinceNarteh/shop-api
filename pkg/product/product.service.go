package product

import (
	"errors"
	"shop_api/pkg/config"
)

func Create(product *Product) {
	config.Database.Db.Create(product)
}

func FindProducts(products *[]Product) {
	config.Database.Db.Find(products)
}

func FindProduct(product *Product, conditions ...interface{}) error {
	result := config.Database.Db.Find(product, conditions...)
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

func SaveProduct(product *Product) {
	config.Database.Db.Save(product)
}

func Delete(product *Product) {
	config.Database.Db.Delete(product)
}
