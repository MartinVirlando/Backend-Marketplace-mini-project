package services

import (
	"backend/models"
	"backend/repositories"
	"errors"
)

type ProductRequest struct {
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryID  uint
	Images      models.StringArray
}

type ProductServiceInterface interface {
	Create(sellerID uint, req ProductRequest) (*models.Product, error)
	GetAll(search string, categoryId uint, page int, limit int) ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Update(id uint, sellerID uint, req ProductRequest) (*models.Product, error)
	Delete(id uint, sellerID uint) error
	GetBySeller(sellerID uint) ([]models.Product, error)
	UpdateStatus(id uint, sellerID uint, status string) (*models.Product, error)
}

type ProductService struct {
	repo repositories.ProductRepositoryInterface
}

func NewProductService(repo repositories.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(sellerID uint, req ProductRequest) (*models.Product, error) {
	//Membuat struct Product Baru
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		Images:      req.Images,
		SellerID:    sellerID,
		Status:      "pending",
	}

	//Simpan ke DB
	err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(product.ID)
}

func (s *ProductService) GetAll(search string, categoryID uint, page int, limit int) ([]models.Product, error) {
	products, err := s.repo.FindAll(search, categoryID, page, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetByID(id uint) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Update(id uint, sellerID uint, req ProductRequest) (*models.Product, error) {
	//Cari Product berdasarkan ID
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if product.SellerID != sellerID {
		return nil, errors.New("you are not allowed to update this product")
	}

	//Update fieldnya
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID
	product.Images = req.Images

	//Simpan ke DB
	err = s.repo.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Delete(id uint, sellerID uint) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if product.SellerID != sellerID {
		return errors.New("you are not allowed to delete this product")
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) GetBySeller(sellerID uint) ([]models.Product, error) {
	products, err := s.repo.FindBySeller(sellerID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) UpdateStatus(id uint, sellerID uint, status string) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if product.SellerID != sellerID {
		return nil, errors.New("you are not allowed to update this product")
	}

	product.Status = status
	err = s.repo.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
