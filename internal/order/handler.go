package order

import (
	"ecommerce/duckyarmy/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	service *orderService1
}

func NewOrderHandler(s *orderService1) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) CheckOut(ctx *gin.Context) {
	userID := auth.GetUserID(ctx)
	if userID == -1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "did not find user"})
		return
	}

	err := h.service.CheckOut(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "order created"})
}

func (h *OrderHandler) GetOrders(ctx *gin.Context) {
	userID := auth.GetUserID(ctx)
	if userID == -1 {
		return
	}

	orders, err := h.service.GetOrders(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := h.service.GetAllOrders(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
