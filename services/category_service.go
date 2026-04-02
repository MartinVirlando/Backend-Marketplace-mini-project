package services

import (
	"backend/models"
	"backend/repositories"
)

type CategoryServiceInterface interface {
	Create(name, icon string) (*models.Category, error)
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	Update(id uint, name, icon string) (*models.Category, error)
	Delete(id uint) error
}

type CategoryService struct {
	repo repositories.CategoryRepositoryInterface
}

func NewCategoryService(repo repositories.CategoryRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(name, icon string) (*models.Category, error) {
	//Membuat struct Category Baru
	category := &models.Category{
		Name: name,
		Icon: icon,
	}

	//Simpan ke DB
	err := s.repo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	categories, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryService) GetByID(id uint) (*models.Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) Update(id uint, name, icon string) (*models.Category, error) {
	//Cari Category berdasarkan ID
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	//Update fieldnya
	category.Name = name
	category.Icon = icon

	//Simpan ke DB
	err = s.repo.Update(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
