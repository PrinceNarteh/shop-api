package repository

import (
	"errors"
	"shop_api/pkg/config"
)

func Create(data interface{}) {
	config.Database.Db.Create(data)
}

func Find(dest interface{}, conditions ...interface{}) error {
	result := config.Database.Db.Find(dest, conditions...)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func FindMany(dest interface{}, conditions ...interface{}) {
	config.Database.Db.Find(dest)
}

func Delete(data interface{}) {
	config.Database.Db.Delete(data)
}

func Save(data interface{}) {
	config.Database.Db.Save(data)
}
