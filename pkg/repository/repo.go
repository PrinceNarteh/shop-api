package repository

import (
	"errors"
	"shop_api/pkg/config"
)

func Create(data interface{}) {
	config.DB.Create(data)
}

func Find(dest interface{}, conditions ...interface{}) error {
	result := config.DB.Find(dest, conditions...)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func FindMany(dest interface{}, conditions ...interface{}) {
	config.DB.Find(dest)
}

func Delete(data interface{}) {
	config.DB.Delete(data)
}

func Save(data interface{}) {
	config.DB.Save(data)
}
