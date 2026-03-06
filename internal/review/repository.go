package review

import (
	"database/sql"
	"fmt"
)

type ReviewRepository interface {
	AddReview(review Review) error
	DeleteReview(commentID, userID int) error
	GetReviewsByProduct(productID int) ([]Review, error)
}

type mysqlReviewRepository struct {
	db *sql.DB
}

func NewMysqlReviewRepository(db *sql.DB) *mysqlReviewRepository {
	return &mysqlReviewRepository{db}
}

func (r *mysqlReviewRepository) AddReview(review Review) error {

	_, err := r.db.Exec(`
		INSERT INTO reviews (product_id, user_id, comment_text, grade, created_at)
		VALUES (?, ?, ?, ?, NOW())
	`,
		review.ProductID,
		review.UserID,
		review.CommentText,
		review.Grade,
	)

	return err
}

func (r *mysqlReviewRepository) DeleteReview(commentID, userID int) error {

	res, err := r.db.Exec(`
		DELETE FROM reviews
		WHERE comment_id = ?
		AND user_id = ?
	`, commentID, userID)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("review not owned by user")
	}

	return nil
}

func (r *mysqlReviewRepository) GetReviewsByProduct(productID int) ([]Review, error) {

	rows, err := r.db.Query(`
		SELECT comment_id, product_id, user_id, comment_text, grade, created_at
		FROM reviews
		WHERE product_id = ?
	`, productID)

	if err != nil {
		return nil, err
	}

	var reviews []Review

	for rows.Next() {

		var r Review

		err := rows.Scan(
			&r.CommentID,
			&r.ProductID,
			&r.UserID,
			&r.CommentText,
			&r.Grade,
			&r.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		reviews = append(reviews, r)
	}

	return reviews, nil
}
