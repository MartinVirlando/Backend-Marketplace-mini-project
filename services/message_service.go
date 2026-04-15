package services

import (
	"backend/models"
	"backend/repositories"
	"log"
)

type MessageServiceInterface interface {
	SendMessage(senderID uint, receiverID uint, productID *uint, message string) (*models.Message, error)
	GetConversations(userID uint) ([]models.Message, error)
	GetMessages(userID uint, otherUserID uint) ([]models.Message, error)
	MarkAsRead(userID uint, senderID uint) error
}

type MessageService struct {
	repo      repositories.MessageRepositoryInterface
	notifRepo repositories.NotificationRepositoryInterface
}

func NewMessageService(repo repositories.MessageRepositoryInterface, notifRepo repositories.NotificationRepositoryInterface) MessageServiceInterface {
	return &MessageService{repo: repo, notifRepo: notifRepo}
}

func (s *MessageService) SendMessage(senderID uint, receiverID uint, productID *uint, message string) (*models.Message, error) {
	msg := &models.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		ProductID:  productID,
		Message:    message,
	}

	err := s.repo.Create(msg)
	if err != nil {
		return nil, err
	}

	existing, _ := s.notifRepo.FindUnreadChat(receiverID, senderID)
	log.Printf("FindUnreadChat result: existing=%v, err=%v", existing, err)
	if existing == nil {
		notif := &models.Notification{
			UserID:      receiverID,
			Title:       "Pesan Baru",
			Message:     "Kamu mendapat pesan baru",
			Type:        "chat",
			IsRead:      false,
			ReferenceID: &senderID,
		}
		s.notifRepo.Create(notif)
	}

	return s.repo.FindByID(msg.ID)
}

func (s *MessageService) GetConversations(userID uint) ([]models.Message, error) {
	messages, err := s.repo.FindConversations(userID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (s *MessageService) GetMessages(userID uint, otherUserID uint) ([]models.Message, error) {
	messages, err := s.repo.FindMessage(userID, otherUserID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (s *MessageService) MarkAsRead(userID uint, senderID uint) error {
	err := s.repo.MarkAsRead(userID, senderID)
	if err != nil {
		return err
	}
	return nil
}
