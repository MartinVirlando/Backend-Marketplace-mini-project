package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Create(product *models.Product) error
	FindByID(id uint) (*models.Product, error)
	FindAll(search string, categoryId uint, page int, limit int) ([]models.Product, error)
	FindBySeller(sellerID uint) ([]models.Product, error)
	FindPending() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Seller").Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) FindAll(search string, categoryId uint, page int, limit int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.Preload("Seller").Preload("Category").Offset((page-1)*limit).Limit(limit).Where("status = ?", "approved")

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if categoryId != 0 {
		query = query.Where("category_id = ?", categoryId)
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) FindBySeller(sellerID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Seller").Preload("Category").Where("seller_id = ?", sellerID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FindPending() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Seller").Preload("Category").Where("status = ?", "pending").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error

}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
