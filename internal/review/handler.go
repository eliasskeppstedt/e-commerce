package review

import (
	"ecommerce/duckyarmy/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReviewHandler struct {
	service *reviewService
}

func NewReviewHandler(s *reviewService) *ReviewHandler {
	return &ReviewHandler{service: s}
}

func (h *ReviewHandler) AddReview(ctx *gin.Context) {
	userID := auth.GetUserID(ctx)
	if userID == -1 {
		return
	}

	var req struct {
		ProductID int    `json:"product_id"`
		Text      string `json:"text"`
		Grade     int    `json:"grade"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.service.AddReview(ctx.Request.Context(), userID, req.ProductID, req.Grade, req.Text)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "review added"})
}

func (h *ReviewHandler) DeleteReview(ctx *gin.Context) {
	userID := auth.GetUserID(ctx)
	if userID == -1 {
		return
	}

	var req struct {
		CommentID int `json:"comment_id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	commentID := req.CommentID

	err := h.service.DeleteReview(ctx.Request.Context(), commentID, userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "review deleted"})
}

func (h *ReviewHandler) GetReviews(ctx *gin.Context) {

	productIDStr := ctx.Query("product_id")
	if productIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product_id required"})
		return
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id"})
		return
	}

	reviews, err := h.service.GetReviews(ctx.Request.Context(), productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}
