package order

import (
	"errors"
	"shop_api/pkg/config"
)

func GetOrders(orders *[]Order) {
	config.Database.Db.Find(orders)
}

func FindProduct(order *Order, conditions ...interface{}) error {
	result := config.Database.Db.Find(order, conditions...)
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

func Save(product *Order) {
	config.Database.Db.Save(product)
}

func Create(order *Order) {
	config.Database.Db.Create(order)
}

func Delete(order *Order) {
	config.Database.Db.Delete(order)
}
