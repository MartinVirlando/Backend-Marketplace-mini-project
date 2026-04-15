package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ReviewRepositoryInterface interface {
	Create(review *models.Review) error
	FindByProductID(productID uint) ([]models.Review, error)
	FindByID(id uint) (*models.Review, error)
	Delete(id uint) error
	FindByUserAndProduct(userID uint, productID uint) (*models.Review, error)
}

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepositoryInterface {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) FindByProductID(productID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Preload("User").Preload("Product").Preload("Product.Seller").Preload("Product.Category").Where("product_id = ?", productID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewRepository) FindByID(id uint) (*models.Review, error) {
	var review models.Review
	err := r.db.Preload("User").Preload("Product").Preload("Product.Seller").Preload("Product.Category").First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) Delete(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}

func (r *ReviewRepository) FindByUserAndProduct(userID uint, productID uint) (*models.Review, error) {
	var review models.Review
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}
