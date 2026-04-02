package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type NotificationRepositoryInterface interface{
	Create(notification *models.Notification) error
	FindByUserID(userID uint) ([]models.Notification, error)
	MarkAsRead(id uint) error
	MarkAllAsRead(userID uint) error
} 

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepositoryInterface {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Create(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *NotificationRepository) FindByUserID(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("user_id = ?", userID).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *NotificationRepository) MarkAsRead(id uint) error {
	return r.db.Model(&models.Notification{}).Where("id = ?", id).Update("is_read", true).Error
}

func (r *NotificationRepository) MarkAllAsRead(userID uint) error {
	return r.db.Model(&models.Notification{}).Where("user_id = ?", userID).Update("is_read", true).Error
}