package services

import (
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(name, email, password, phone, role string) (*models.User, error)
	Login(email, password string) (string, error)
	GetMe(id uint) (*models.User, error)
	UpdateProfile(id uint, name, phone string) (*models.User, error)
}

type AuthService struct {
	repo repositories.UserRepositoryInterface
}

func NewAuthService(repo repositories.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(name, email, password, phone, role string) (*models.User, error) {
	//Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	//Membuat struct User Baru
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Phone:    phone,
		Role:     role,
	}

	//Simpan ke DB
	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	//Cari User berdasarkan Email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	//Cek Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	//Generate Token
	token, err := utils.GenerateToken(fmt.Sprintf("%d", user.ID), user.Role)
	if err != nil {
		return "", err
	}

	//Return Token
	return token, nil
}

func (s *AuthService) GetMe(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) UpdateProfile(id uint, name, phone string) (*models.User, error) {
	//Cari User berdasarkan ID
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	//Update
	user.Name = name
	user.Phone = phone
	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
