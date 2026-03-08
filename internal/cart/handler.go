package cart

import (
	"ecommerce/duckyarmy/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *cartService1
}

func NewCartHandler(s *cartService1) *CartHandler {
	return &CartHandler{s}
}

func (h *CartHandler) AddItem(ctx *gin.Context) {
	claimsValue, exists := ctx.Get("auth_token")

	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "user not authenticated"})
		return
	}

	claims := claimsValue.(*auth.Claims)
	userID := claims.UserID

	var req AddToCartRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	productID := req.ProductID
	// quantity är inte satt än i frontend, atm lägg ändast till 1 per tryck
	quantity := 1 // req.Quantity

	err := h.service.AddItem(ctx.Request.Context(), userID, productID, quantity)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not add item"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "product added successfully"})
}
