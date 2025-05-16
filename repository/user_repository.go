package repository

import "pre_tpa_web/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(email string, updatedUser *model.User) error
	DeleteUser(email string) error
}