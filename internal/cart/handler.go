package cart

import (
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *cartService1
}

func NewCartHandler(s *cartService1) *CartHandler {
	return &CartHandler{s}
}

func (h *CartHandler) AddItem(ctx *gin.Context) {
	// fixa med auth
}
