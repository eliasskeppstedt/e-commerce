package order

import (
	"ecommerce/duckyarmy/internal/auth"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *orderService1
}

func NewOrderHandler(s *orderService1) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) CheckOut(ctx *gin.Context) {

	claimsValue, exists := ctx.Get("auth_token")
	if !exists {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, ok := claimsValue.(*auth.Claims)
	if !ok {
		ctx.JSON(500, gin.H{"error": "invalid auth token"})
		return
	}

	err := h.service.CheckOut(claims.UserID)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "order created"})
}

func (h *OrderHandler) GetOrders(ctx *gin.Context) {

	claimsValue, _ := ctx.Get("auth_token")
	claims := claimsValue.(*auth.Claims)

	orders, err := h.service.GetOrders(claims.UserID)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}
