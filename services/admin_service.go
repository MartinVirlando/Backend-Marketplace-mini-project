package services

import (
	"backend/models"
	"backend/repositories"
)

type DashboardStats struct {
	TotalTransactions int
	TotalUsers        int
	TotalProducts     int
}

type AdminServiceInterface interface {
	GetDashboardStats() (*DashboardStats, error)
	GetPendingProducts() ([]models.Product, error)
	ApproveProduct(id uint) error
	RejectProduct(id uint) error
	GetUsers() ([]models.User, error)
	DeleteUser(id uint) error
	ApproveAllProducts() error
}

type AdminService struct {
	productRepo repositories.ProductRepositoryInterface
	userRepo    repositories.UserRepositoryInterface
	orderRepo   repositories.OrderRepositoryInterface
}

func NewAdminService(
	productRepo repositories.ProductRepositoryInterface,
	userRepo repositories.UserRepositoryInterface,
	orderRepo repositories.OrderRepositoryInterface,
) AdminServiceInterface {
	return &AdminService{
		productRepo: productRepo,
		userRepo:    userRepo,
		orderRepo:   orderRepo,
	}
}

func (s *AdminService) GetDashboardStats() (*DashboardStats, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	products, err := s.productRepo.FindAll("", 0, 1, 1000)
	if err != nil {
		return nil, err
	}

	orders, err := s.orderRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &DashboardStats{
		TotalTransactions: len(orders),
		TotalUsers:        len(users),
		TotalProducts:     len(products),
	}, nil
}

func (s *AdminService) GetPendingProducts() ([]models.Product, error) {
	products, err := s.productRepo.FindPending()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *AdminService) ApproveProduct(id uint) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}
	product.Status = "approved"
	return s.productRepo.Update(product)
}

func (s *AdminService) RejectProduct(id uint) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}
	product.Status = "rejected"
	return s.productRepo.Update(product)
}

func (s *AdminService) ApproveAllProducts() error {
	products, err := s.productRepo.FindPending()
	if err != nil {
		return err
	}
	for _, product := range products {
		product.Status = "approved"
		err = s.productRepo.Update(&product)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *AdminService) GetUsers() ([]models.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *AdminService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
