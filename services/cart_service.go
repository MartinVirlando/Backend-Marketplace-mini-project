package services

import (
	"backend/models"
	"backend/repositories"
	"errors"
)

type CartServiceInterface interface {
	AddItem(userID uint, productID uint, quantity int) (*models.CartItem, error)
	GetCart(userID uint) ([]models.CartItem, error)
	UpdateCart(userID uint, cartID uint, quantity int) (*models.CartItem, error)
	DeleteCart(userID uint, cartID uint) error
	ClearCart(userID uint) error
}

type CartService struct {
	repo repositories.CartRepositoryInterface
}

func NewCartService(repo repositories.CartRepositoryInterface) CartServiceInterface {
	return &CartService{repo: repo}
}

func (s *CartService) AddItem(userID uint, productID uint, quantity int) (*models.CartItem, error) {
	cart := &models.CartItem{
		Quantity:  quantity,
		UserID:    userID,
		ProductID: productID,
	}
	err := s.repo.AddItem(cart)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(cart.ID)
}

func (s *CartService) GetCart(userID uint) ([]models.CartItem, error) {
	carts, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (s *CartService) UpdateCart(userID uint, cartID uint, quantity int) (*models.CartItem, error) {
	cart, err := s.repo.FindByID(cartID)
	if err != nil {
		return nil, err
	}

	if cart.UserID != userID {
		return nil, errors.New("you are not allowed to update this cart")
	}

	cart.Quantity = quantity

	err = s.repo.Update(cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *CartService) DeleteCart(userID uint, cartID uint) error {
	cart, err := s.repo.FindByID(cartID)
	if err != nil {
		return err
	}

	if cart.UserID != userID {
		return errors.New("you are not allowed to delete this cart")
	}

	err = s.repo.Delete(cartID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) ClearCart(userID uint) error {
	err := s.repo.ClearByUserID(userID)
	if err != nil {
		return err
	}
	return nil
}
