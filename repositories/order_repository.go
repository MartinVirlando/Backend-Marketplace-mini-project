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
	UpdateCartItem(cartItem *models.CartItem) error
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
	err := r.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	r.db.First(&order.User, order.UserID)

	var items []models.CartItem
	r.db.Where("order_id = ?", id).Find(&items)

	for i := range items {
		r.db.First(&items[i].Product, items[i].ProductID)
		r.db.First(&items[i].Product.Seller, items[i].Product.SellerID)
		r.db.First(&items[i].Product.Category, items[i].Product.CategoryID)
		r.db.First(&items[i].User, items[i].UserID)
	}

	order.Items = items
	return &order, nil
}

func (r *OrderRepository) FindByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	r.db.Where("user_id = ?", userID).Find(&orders)

	for i := range orders {
		r.db.First(&orders[i].User, orders[i].UserID)

		var items []models.CartItem
		r.db.Where("order_id = ?", orders[i].ID).Find(&items)

		for j := range items {
			r.db.First(&items[j].Product, items[j].ProductID)
			r.db.First(&items[j].Product.Seller, items[j].Product.SellerID)
			r.db.First(&items[j].Product.Category, items[j].Product.CategoryID)
			r.db.First(&items[j].User, items[j].UserID)
		}

		orders[i].Items = items
	}

	return orders, nil
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	r.db.Find(&orders)

	for i := range orders {
		r.db.First(&orders[i].User, orders[i].UserID)

		var items []models.CartItem
		r.db.Where("order_id = ?", orders[i].ID).Find(&items)

		for j := range items {
			r.db.First(&items[j].Product, items[j].ProductID)
			r.db.First(&items[j].Product.Seller, items[j].Product.SellerID)
			r.db.First(&items[j].Product.Category, items[j].Product.CategoryID)
			r.db.First(&items[j].User, items[j].UserID)
		}

		orders[i].Items = items
	}

	return orders, nil
}

func (r *OrderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) UpdateCartItem(cartItem *models.CartItem) error {
	return r.db.Save(cartItem).Error
}
