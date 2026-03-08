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
	userID := auth.GetUserID(ctx)

	if userID == -1 {
		return
	}

	var req AddToCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func (h *CartHandler) RequestCartItems(ctx *gin.Context) {
	userID := auth.GetUserID(ctx)
	if userID == -1 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "user is unauthorized"})
		return
	}

	cartItems, err := h.service.RequestCartItems(ctx.Request.Context(), userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cartItems)
}
