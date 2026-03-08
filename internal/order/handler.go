package order

import (
	"fmt"
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
	fmt.Println("OrderHandler CheckOut: hårdkodat userID = 1")
	err := h.service.CheckOut(ctx.Request.Context(), 1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Redirect(http.StatusSeeOther, "/cartPage")
		return
	}
	ctx.JSON(200, gin.H{"message": "order created"})
	ctx.Redirect(http.StatusSeeOther, "/homePage")
}
