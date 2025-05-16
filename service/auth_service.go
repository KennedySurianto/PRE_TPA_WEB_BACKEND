package service

import (
	"errors"
	"fmt"
	"pre_tpa_web/model"
	"pre_tpa_web/repository"
	"pre_tpa_web/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

func (a *AuthServiceImpl) Register(name, email, password string) (string, error) {
	// Cek apakah email sudah digunakan
	existing, err := a.userRepo.GetUserByEmail(email)
	if err == nil && existing != nil {
		// found a user, return error
		return "", errors.New("email already registered")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		// some other DB error (e.g. connection issue)
		return "", err
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	createErr := a.userRepo.CreateUser(user)
	if err != nil {
		return "", createErr
	}

	return "User registered successfully", nil
}

func (a *AuthServiceImpl) Login(email, password string) (string, error) {
	fmt.Println("from login: email -> : " + email)

	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
