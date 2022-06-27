package order

import (
	"errors"
	"shop_api/pkg/config"
)

func GetOrders(orders *[]Order) {
	config.DB.Find(orders)
}

func FindProduct(order *Order, conditions ...interface{}) error {
	result := config.DB.Find(order, conditions...)
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

func Save(product *Order) {
	config.DB.Save(product)
}

func Create(order *Order) {
	config.DB.Create(order)
}

func Delete(order *Order) {
	config.DB.Delete(order)
}
