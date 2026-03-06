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
	claimsValue, exists := ctx.Get("auth_token")
	if !exists {
		return
	}
	claims := claimsValue.(*auth.Claims)
	userID := claims.UserID

	var req struct {
		ProductID int    `json:"product_id"`
		Text      string `json:"text"`
		Grade     int    `json:"grade"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.service.AddReview(userID, req.ProductID, req.Grade, req.Text)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "review added"})
}

func (h *ReviewHandler) DeleteReview(ctx *gin.Context) {

	claimsValue, exists := ctx.Get("auth_token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims := claimsValue.(*auth.Claims)
	userID := claims.UserID

	idStr := ctx.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	err = h.service.DeleteReview(commentID, userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "review deleted"})
}

func (h *ReviewHandler) GetReviews(ctx *gin.Context) {

	productIDStr := ctx.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)

	if err != nil {
		return
	}

	reviews, err := h.service.GetReviews(productID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}
