package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type MessageRepositoryInterface interface {
	Create(message *models.Message) error
	FindByID(id uint) (*models.Message, error)
	FindConversations(userID uint) ([]models.Message, error)
	FindMessage(userID uint, otherUserID uint) ([]models.Message, error)
	MarkAsRead(userID uint, senderID uint) error
}

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepositoryInterface {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepository) FindConversations(userID uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Preload("Sender").Preload("Receiver").Preload("Product").Preload("Product.Seller").Preload("Product.Category").Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Order("created_at DESC").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *MessageRepository) FindMessage(userID uint, otherUserID uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Preload("Sender").Preload("Receiver").Preload("Product").Preload("Product.Seller").Preload("Product.Category").
		Where(
			"(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			userID, otherUserID, otherUserID, userID,
		).
		Order("created_at ASC").
		Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *MessageRepository) MarkAsRead(userID uint, senderID uint) error {
	return r.db.Model(&models.Message{}).Where("receiver_id = ? AND sender_id = ? AND is_read = ?", userID, senderID, false).Update("is_read", true).Error
}

func (r *MessageRepository) FindByID(id uint) (*models.Message, error) {
	var message models.Message
	err := r.db.Preload("Sender").Preload("Receiver").Preload("Product").Preload("Product.Seller").Preload("Product.Category").First(&message, id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}
