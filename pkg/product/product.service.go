package product

import "shop_api/pkg/config"

func Create(product *Product) {
	config.Database.Db.Create(product)
}