package service

import "pre_tpa_web/model"

type UserService interface {
	CreateUser(name, email, password string) (string, error)
	GetAllUsers() ([]model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(email, name, password string) error
	DeleteUser(email string) error
}
