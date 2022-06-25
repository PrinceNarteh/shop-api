package user

import (
	"shop_api/pkg/config"
)

func FindUser(userId uint, user *User) {
	config.Database.Db.Find(&user, userId)
}

func FindUserByEmail(email string, user *User) {
	config.Database.Db.Find(&user, "email = ?", email)
}

func FindUsers(users *[]User) {
	config.Database.Db.Find(&users)
}
