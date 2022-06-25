package user

import (
	"errors"
	"shop_api/pkg/config"
)

func FindUser(user *User, conditions ...interface{}) error {
	config.Database.Db.Find(&user, conditions...)
	if user.ID == 0 {
		return errors.New("user not found")
	}
	return nil
}

func FindUsers(users *[]User) {
	config.Database.Db.Find(&users)
}

func SaveUser(user *User) {
	config.Database.Db.Save(user)
}
