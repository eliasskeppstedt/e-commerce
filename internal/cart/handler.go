package cart

import (
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *CartService1
}

func NewCartHandler(s *CartService1) *CartHandler {
	return &CartHandler{s}
}

func (h *CartHandler) CartItemAdd(ctx *gin.Context) {
}

func (h *CartHandler) CartItemBulkAdd(ctx *gin.Context) {
}

func (h *CartHandler) CartItemRemove(ctx *gin.Context) {
}

func (h *CartHandler) CartItemBulkRemove(ctx *gin.Context) {
}
