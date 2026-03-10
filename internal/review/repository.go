package review

import (
	"context"
	"database/sql"
	"fmt"
)

type ReviewRepository interface {
	AddReview(ctx context.Context, tx *sql.Tx, review Review) error
	DeleteReview(ctx context.Context, tx *sql.Tx, commentID int) error
	GetReviewsByProduct(ctx context.Context, tx *sql.Tx, productID int) ([]Review, error)
	GetReviewByID(ctx context.Context, tx *sql.Tx, commentID int) (*Review, error)
	HasUserReviewedProduct(ctx context.Context, tx *sql.Tx, userID, productID int) (bool, error)
}

type mysqlReviewRepository struct {
	db *sql.DB
}

func NewMysqlReviewRepository(db *sql.DB) *mysqlReviewRepository {
	return &mysqlReviewRepository{db}
}

func (r *mysqlReviewRepository) AddReview(ctx context.Context, tx *sql.Tx, review Review) error {

	_, err := tx.ExecContext(
		ctx,
		`INSERT INTO reviews (product_id, user_id, comment_text, grade)
		VALUES (?, ?, ?, ?)`,
		review.ProductID,
		review.UserID,
		review.CommentText,
		review.Grade,
	)

	return err
}

func (r *mysqlReviewRepository) DeleteReview(ctx context.Context, tx *sql.Tx, commentID int) error {

	_, err := tx.ExecContext(
		ctx,
		`DELETE FROM reviews
		WHERE comment_id = ?`,
		commentID,
	)

	return err
}

func (r *mysqlReviewRepository) GetReviewByID(ctx context.Context, tx *sql.Tx, commentID int) (*Review, error) {
	var review Review
	err := tx.QueryRowContext(
		ctx,
		`SELECT comment_id, product_id, user_id, comment_text, grade, created_at
		 FROM reviews
		 WHERE comment_id = ?`,
		commentID,
	).Scan(
		&review.CommentID,
		&review.ProductID,
		&review.UserID,
		&review.CommentText,
		&review.Grade,
		&review.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("review not found")
		}
		return nil, err
	}
	return &review, nil
}

func (r *mysqlReviewRepository) HasUserReviewedProduct(ctx context.Context, tx *sql.Tx, userID, productID int) (bool, error) {
	var count int
	err := tx.QueryRowContext(
		ctx,
		`SELECT COUNT(*) FROM reviews WHERE user_id = ? AND product_id = ?`,
		userID,
		productID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *mysqlReviewRepository) GetReviewsByProduct(ctx context.Context, tx *sql.Tx, productID int) ([]Review, error) {

	rows, err := tx.QueryContext(
		ctx,
		`SELECT 
			r.comment_id,
			r.product_id,
			r.user_id,
			u.username,
			r.comment_text,
			r.grade,
			r.created_at
		FROM reviews r
		JOIN users u ON r.user_id = u.user_id
		WHERE r.product_id = ?`,
		productID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var reviews []Review

	for rows.Next() {

		var review Review

		err := rows.Scan(
			&review.CommentID,
			&review.ProductID,
			&review.UserID,
			&review.Username,
			&review.CommentText,
			&review.Grade,
			&review.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}
