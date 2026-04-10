package services

import (
	"backend/models"
	"backend/repositories"
	"errors"
)

type OrderServiceInterface interface {
	CreateOrder(userID uint, cartItems []models.CartItem, shippingAddress, city, province, postalCode string) (*models.Order, error)
	GetOrders(userID uint) ([]models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	CancelOrder(id uint, userID uint) error
}

type OrderService struct {
	repo repositories.OrderRepositoryInterface
}

func NewOrderService(repo repositories.OrderRepositoryInterface) OrderServiceInterface {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(userID uint, cartItems []models.CartItem, shippingAddress, city, province, postalCode string) (*models.Order, error) {
	//Total Price
	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += item.Product.Price * float64(item.Quantity)
	}
	//Membuat struct Order Baru
	order := &models.Order{
		UserID:          userID,
		ShippingAddress: shippingAddress,
		City:            city,
		Province:        province,
		PostalCode:      postalCode,
		TotalPrice:      totalPrice,
	}

	//Simpan ke DB
	err := s.repo.Create(order)
	if err != nil {
		return nil, err
	}

	for i := range cartItems {
		cartItems[i].OrderID = &order.ID
		err = s.repo.UpdateCartItem(&cartItems[i])
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (s *OrderService) GetOrders(userID uint) ([]models.Order, error) {
	orders, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) CancelOrder(id uint, userID uint) error {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	//Check Ownership
	if order.UserID != userID {
		return errors.New("you are not allowed to cancel this order")
	}

	//Check Apakah status Pending
	if order.Status != "pending" {
		return errors.New("you can't cancel this order")
	}

	//Update Status
	order.Status = "cancelled"
	err = s.repo.Update(order)
	if err != nil {
		return err
	}

	return nil
}
