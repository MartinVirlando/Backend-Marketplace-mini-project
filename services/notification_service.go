package services

import (
	"backend/models"
	"backend/repositories"
)

type NotificationServiceInterface interface {
	CreateNotification(userID uint, title, message, notifType string) error
	GetNotifications(userID uint) ([]models.Notification, error)
	MarkAsRead(id uint) error
	MarkAllAsRead(userID uint) error
	DeleteNotification(id uint) error
}

type NotificationService struct {
	repo repositories.NotificationRepositoryInterface
}

func NewNotificationService(repo repositories.NotificationRepositoryInterface) NotificationServiceInterface {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) CreateNotification(userID uint, title, message, notifType string) error {
	notification := &models.Notification{
		UserID:  userID,
		Title:   title,
		Message: message,
		Type:    notifType,
	}
	err := s.repo.Create(notification)
	if err != nil {
		return err
	}
	return nil
}

func (s *NotificationService) GetNotifications(userID uint) ([]models.Notification, error) {
	notifications, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (s *NotificationService) MarkAsRead(id uint) error {
	err := s.repo.MarkAsRead(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *NotificationService) MarkAllAsRead(userID uint) error {
	err := s.repo.MarkAllAsRead(userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *NotificationService) DeleteNotification(id uint) error {
	return s.repo.Delete(id)
}
