package user

import (
	"errors"
	"shop_api/pkg/config"
)

func FindUser(user *User, conditions ...interface{}) error {
	result := config.DB.Find(&user, conditions...)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func FindUsers(users *[]User) {
	config.DB.Find(users)
}

func Delete(user *User) {
	config.DB.Delete(user)
}

func SaveUser(user *User) {
	config.DB.Save(user)
}
