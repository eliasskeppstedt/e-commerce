package cart

import (
	"ecommerce/duckyarmy/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
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

	// quantity är inte satt än i frontend, atm lägg ändast till 1 per tryck
	err := h.service.AddItem(ctx.Request.Context(), userID, req.ProductID, 1) //quantity)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func (h *CartHandler) RemoveItem(ctx *gin.Context) {
	userID := auth.GetUserID(ctx)

	if userID == -1 {
		return
	}

	var req AddToCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.RemoveItem(ctx.Request.Context(), userID, req.ProductID, 1) // req.Quantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "cart item successfully deleted"})
}
