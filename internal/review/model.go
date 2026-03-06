package review

import "time"

type Review struct {
	CommentID   int       `json:"comment_id"`
	ProductID   int       `json:"product_id"`
	UserID      int       `json:"user_id"`
	CommentText string    `json:"comment_text"`
	Grade       int       `json:"grade"`
	CreatedAt   time.Time `json:"created_at"`
}
