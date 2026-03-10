package review

import (
	"context"
	"ecommerce/duckyarmy/internal/transaction"
	"errors"
)

type ReviewService interface {
	AddReview(ctx context.Context, userID, productID, grade int, text string) error
	DeleteReview(ctx context.Context, commentID, userID int) error
	GetReviews(ctx context.Context, productID int) ([]Review, error)
}

type reviewService struct {
	tm   transaction.TxManager
	repo ReviewRepository
}

func NewReviewService(tm transaction.TxManager, repo ReviewRepository) *reviewService {
	return &reviewService{tm: tm, repo: repo}
}

func (s *reviewService) AddReview(ctx context.Context, userID, productID, grade int, text string) error {
	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if grade < 1 || grade > 5 {
		return errors.New("grade must be between 1 and 5")
	}

	hasReviewed, err := s.repo.HasUserReviewedProduct(ctx, tx, userID, productID)
	if err != nil {
		return err
	}
	if hasReviewed {
		return errors.New("user has already reviewed this product")
	}

	review := Review{
		UserID:      userID,
		ProductID:   productID,
		CommentText: text,
		Grade:       grade,
	}

	err = s.repo.AddReview(ctx, tx, review)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *reviewService) DeleteReview(ctx context.Context, commentID, userID int) error {
	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	review, err := s.repo.GetReviewByID(ctx, tx, commentID)
	if err != nil {
		return err
	}

	if review.UserID != userID {
		return errors.New("review not owned by user")
	}

	err = s.repo.DeleteReview(ctx, tx, commentID)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *reviewService) GetReviews(ctx context.Context, productID int) ([]Review, error) {
	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	reviews, err := s.repo.GetReviewsByProduct(ctx, tx, productID)
	if err != nil {
		return nil, err
	}
	return reviews, tx.Commit()
}
