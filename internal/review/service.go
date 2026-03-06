package review

import "errors"

type ReviewService interface {
	AddReview(userID, productID, grade int, text string) error
	DeleteReview(commentID, userID int) error
	GetReviews(productID int) ([]Review, error)
}

type reviewService struct {
	repo ReviewRepository
}

func NewReviewService(repo ReviewRepository) *reviewService {
	return &reviewService{repo}
}

func (s *reviewService) AddReview(userID, productID, grade int, text string) error {

	if grade < 1 || grade > 5 {
		return errors.New("grade must be between 1 and 5")
	}

	review := Review{
		UserID:      userID,
		ProductID:   productID,
		CommentText: text,
		Grade:       grade,
	}

	return s.repo.AddReview(review)
}

func (s *reviewService) DeleteReview(commentID, userID int) error {
	return s.repo.DeleteReview(commentID, userID)
}

func (s *reviewService) GetReviews(productID int) ([]Review, error) {
	return s.repo.GetReviewsByProduct(productID)
}
