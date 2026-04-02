package services

import (
	"backend/models"
	"backend/repositories"
	"errors"
)

type ReviewServiceInterface interface {
	CreateReview(userID uint, productID uint, rating int, comment string) (*models.Review, error)
	GetReviews(productID uint) ([]models.Review, error)
	DeleteReview(id uint, userID uint) error
}

type ReviewService struct {
	repo repositories.ReviewRepositoryInterface
}

func NewReviewService(repo repositories.ReviewRepositoryInterface) ReviewServiceInterface {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) CreateReview(userID uint, productID uint, rating int, comment string) (*models.Review, error) {
	review := &models.Review{
		UserID:    userID,
		ProductID: productID,
		Rating:    rating,
		Comment:   comment,
	}
	err := s.repo.Create(review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s *ReviewService) GetReviews(productID uint) ([]models.Review, error) {
	reviews, err := s.repo.FindByProductID(productID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s *ReviewService) DeleteReview(id uint, userID uint) error {
	review, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if review.UserID != userID {
		return errors.New("you are not allowed to delete this review")
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
