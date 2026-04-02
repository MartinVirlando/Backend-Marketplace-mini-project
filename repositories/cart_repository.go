package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type CartRepositoryInterface interface {
	AddItem(cart *models.CartItem) error
	FindByUserID(userID uint) ([]models.CartItem, error)
	FindByID(id uint) (*models.CartItem, error)
	Update(cart *models.CartItem) error
	Delete(id uint) error
	ClearByUserID(userID uint) error
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepositoryInterface {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddItem(cart *models.CartItem) error {
	return r.db.Create(cart).Error
}

func (r *CartRepository) FindByUserID(userID uint) ([]models.CartItem, error) {
	var carts []models.CartItem
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *CartRepository) FindByID(id uint) (*models.CartItem, error) {
	var cart models.CartItem
	err := r.db.First(&cart, id).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) Update(cart *models.CartItem) error {
	return r.db.Save(cart).Error
}

func (r *CartRepository) Delete(id uint) error {
	return r.db.Delete(&models.CartItem{}, id).Error
}

func (r *CartRepository) ClearByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error
}
