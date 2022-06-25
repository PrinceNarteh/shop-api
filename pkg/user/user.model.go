package user

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func CreateUserResponse(userModel *User) UserResponse {
	return UserResponse{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
	}
}
