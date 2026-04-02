package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	Create(order *models.Order) error
	FindByID(id uint) (*models.Order, error)
	FindByUserID(userID uint) ([]models.Order, error)
	Update(order *models.Order) error
	GetAll() ([]models.Order, error)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepositoryInterface {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items.Product").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items.Product").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

